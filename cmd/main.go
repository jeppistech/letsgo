package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "letsgo",
	Short: "An awesome go application!",
	Long:  `This is an example application written in Go`,
}
