package main

import (
	"fmt"

	"github.com/nexustar/quickvm"
	"github.com/spf13/cobra"
)

var (
	flagPublish []string
	ropt        quickvm.RunOpt
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <vm name>",
	Short: "run virtual machine.",
	Long:  `run virtual machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		var err error
		ropt.PortForward, err = quickvm.ParserOptPublish(flagPublish)
		cobra.CheckErr(err)
		fmt.Println("port forward table:")
		fmt.Println("protocol\thost port\tport")
		for _, p := range ropt.PortForward {
			fmt.Printf("%s\t\t%d\t\t%d\n", p.Protocol, p.HostPort, p.Port)
		}

		ropt.Name = args[0]
		ropt.AdditionalArgs = args[1:]

		cobra.CheckErr(quickvm.Run(ropt))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntVarP(&ropt.Cpu, "cpu", "c", 0, "cpu core of vm")
	runCmd.Flags().StringVarP(&ropt.Memory, "mem", "m", "4G", "memory of vm")
	runCmd.Flags().StringSliceVarP(&flagPublish, "publish", "p", nil, "publish ports inside vm to host")
}
