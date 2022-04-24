package main

import (
	"fmt"

	"github.com/nexustar/quickvm"
	"github.com/spf13/cobra"
)

var flagPublish []string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <vm name>",
	Short: "run virtual machine.",
	Long:  `run virtual machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		portsMap, err := quickvm.ParserOptPublish(flagPublish)
		cobra.CheckErr(err)
		fmt.Println("port forward table:")
		fmt.Println("protocol\thost port\tport")
		for _,p := range portsMap{
			fmt.Printf("%s\t\t%d\t\t%d\n", p.Protocol, p.HostPort, p.Port)
		}
		ropt := quickvm.RunOpt{
			Name:        args[0],
			PortForward: portsMap,
		}
		cobra.CheckErr(quickvm.Run(ropt))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringSliceVarP(&flagPublish, "publish", "p", nil, "publish ports inside vm to host")
}
