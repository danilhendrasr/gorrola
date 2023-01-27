package gorrola

import (
	"github.com/danilhendrasr/gorrola/pkg/gorrola"
	"github.com/spf13/cobra"
)

var runCmd = cobra.Command{
	Use:   "run",
	Short: "Run the load balancer",
	Run: func(cmd *cobra.Command, args []string) {
		gorrola.Run()
	},
}

func init() {
	rootCmd.AddCommand(&runCmd)
}
