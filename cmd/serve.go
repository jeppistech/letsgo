package cmd

import (
	"github.com/spf13/cobra"

	"../pkg/server"
)

var (
	defaultPort = "8080"

	port = ""
)

func init() {
	RootCmd.AddCommand(cmdServe)

	cmdServe.Flags().StringVarP(&port, "port", "p", defaultPort, "the port this application will listen on")
}

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Serve a web application",
	Long:  `Serve a REST web application for saving weather data`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve(port)
	},
}
