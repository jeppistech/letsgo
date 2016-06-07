package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"../pkg/version"
)

func init() {
	RootCmd.AddCommand(cmdVersion)
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of this app",
	Long:  `All software has versions!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version %s of course", version.Version)
	},
}
