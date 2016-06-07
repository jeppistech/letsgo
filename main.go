package main

import (
	"./cmd"
	"fmt"
	"os"
)

func main() {

	// Just invokes cobra files in cmd/
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
