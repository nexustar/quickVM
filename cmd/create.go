package main

import (
	"github.com/nexustar/quickvm"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <vm name>",
	Short: "create virtual machine.",
	Long:  `create virtual machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		copt := quickvm.CreateOpt{
			Name: args[0],
		}
		quickvm.Create(copt)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
