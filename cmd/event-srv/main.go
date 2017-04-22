package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version     string
	versionFlag bool
)

func main() {
	versionString := "Event Service v" + version
	cobra.OnInitialize(func() {
		if versionFlag {
			fmt.Println(versionString)
			os.Exit(0)
		}
	})

	var rootCmd = &cobra.Command{
		Use:   "event-srv",
		Short: "Event Service",
		Long:  versionString,
		Run:   Serve,
	}

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print application version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
