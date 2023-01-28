package gorrola

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/danilhendrasr/gorrola/pkg/gorrola"
	"github.com/spf13/cobra"
)

type Flag uint

const (
	backends Flag = iota
	port
)

func (f Flag) String() string {
	switch f {
	case backends:
		return "backends"
	case port:
		return "port"
	}

	return "unknown"
}

var backendsFlag string
var portFlag uint

var runCmd = cobra.Command{
	Use:   "run",
	Short: "Run the load balancer",
	Run: func(cmd *cobra.Command, args []string) {
		backendUrls := strings.Split(backendsFlag, ",")
		for _, u := range backendUrls {
			if _, err := url.ParseRequestURI(u); err != nil {
				fmt.Fprintf(os.Stderr, "%s is not a valid URL.\n", u)
				os.Exit(1)
			}
		}

		gorrola.Run(backendUrls, portFlag)
	},
}

func init() {
	runCmd.Flags().StringVar(&backendsFlag, backends.String(), "", "comma-separated string backends URLs (required)")
	runCmd.Flags().UintVarP(&portFlag, port.String(), "p", 3000, "the port where the load balancer will be served at")
	runCmd.MarkFlagRequired(backends.String())

	rootCmd.AddCommand(&runCmd)
}
