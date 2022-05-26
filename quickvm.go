package quickvm

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
)

var dataDir string

func init() {
	dataDir = filepath.Join(xdg.DataHome, "quickVM")
}

type CreateOpt struct {
	Name string
}

type RunOpt struct {
	Name           string
	PortForward    []PortForward
	AdditionalArgs []string
}

type PortForward struct {
	Protocol string
	Port     int
	HostPort int
}

func Create(opt CreateOpt) error {
	vmDir := filepath.Join(dataDir, "vms", opt.Name)
	err := os.MkdirAll(vmDir, 0700)
	if err != nil {
		return err
	}
	// TBD
	return nil
}

func Run(opt RunOpt) error {
	vmDir := filepath.Join(dataDir, "vms", opt.Name)

	var arch string
	if runtime.GOARCH == "amd64" {
		arch = "x86_64"
	} else {
		arch = "aarch64"
	}

	var portfwd string
	for _, p := range opt.PortForward {
		portfwd = portfwd + fmt.Sprintf(",hostfwd=%s::%d-:%d", p.Protocol, p.HostPort, p.Port)
	}

	binPath, _ := exec.LookPath("qemu-system-" + arch)

	args := []string{
		"-machine",
		"accel=kvm,type=q35",
		"-cpu", "host",
		"-m", "4G",
		"-nographic",
		"-device", "virtio-net-pci,netdev=net0",
		"-netdev", "user,id=net0" + portfwd,
		"-drive", "if=virtio,format=qcow2,file=" + filepath.Join(vmDir, opt.Name+".qcow2"),
		"-drive", "if=virtio,format=raw,file=" + filepath.Join(vmDir, "seed.img"),
	}
	args = append(args, opt.AdditionalArgs...)

	cmd := exec.Command(binPath, args...)
	cmd.Env = os.Environ()
	err := cmd.Run()
	return err
}

func ParserOptPublish(ports []string) ([]PortForward, error) {
	portfwds := make([]PortForward, 0)
	for _, p := range ports {
		var fwd PortForward
		var err error
		split := strings.Split(p, "/")
		switch len(split) {
		case 1:
			fwd.Protocol = "tcp"
		case 2:
			fwd.Protocol = split[1]
		default:
			return nil, fmt.Errorf("cannot parser %s", p)
		}
		split = strings.Split(split[0], ":")
		switch len(split) {
		case 1:
			fwd.Port, err = strconv.Atoi(split[0])
			if err != nil {
				return nil, err
			}
			l, err := net.Listen(fwd.Protocol, fmt.Sprintf("localhost:%d", fwd.Port))
			if err == nil {
				l.Close()
				fwd.HostPort = fwd.Port
			} else {
				l, err = net.Listen(fwd.Protocol, fmt.Sprintf("localhost:0"))
				if err != nil {
					return nil, err
				}
				l.Close()
				fwd.HostPort = l.Addr().(*net.TCPAddr).Port
			}
		case 2:
			fwd.HostPort, err = strconv.Atoi(split[0])
			if err != nil {
				return nil, err
			}
			fwd.Port, err = strconv.Atoi(split[1])
			if err != nil {
				return nil, err
			}
			l, err := net.Listen(fwd.Protocol, fmt.Sprintf("localhost:%d", fwd.HostPort))
			if err != nil {
				return nil, err
			}
			l.Close()
		default:
			return nil, fmt.Errorf("cannot parser %s", p)
		}

		portfwds = append(portfwds, fwd)
	}
	return portfwds, nil
}
