package gorrola

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "gorrola",
	Short: "gorrola - simple Round Robin load balancer written in Go",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error executing gorrola: %s.", err)
		os.Exit(1)
	}
}
