package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/pbar1/swim/pkg/mirror"
)

const (
	pool    = 10
	timeout = 5 * time.Second
)

// getCmd represents the create command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Download all times available for a given parameter space",
	Long:  `Download all times available for a given parameter space`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "d3" || args[0] == "D3" {
			mirror.MirrorNCAAD3(pool, timeout)
		}
	},
}

func init() {
	rootCmd.AddCommand(mirrorCmd)
}
