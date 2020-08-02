package cli

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/pbar1/swim/pkg/mirror"
)

const (
	pool    = 10
	timeout = 6 * time.Second
)

// getCmd represents the create command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Download all times available for a given parameter space",
	Long:  `Download all times available for a given parameter space`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("must give argument: d1|d2|d3")
		}
		switch strings.ToLower(args[0]) {
		case "d1":
			mirror.MirrorNCAAD1(pool, timeout)
		case "d2":
			mirror.MirrorNCAAD2(pool, timeout)
		case "d3":
			mirror.MirrorNCAAD3(pool, timeout)
		default:
			log.Fatalf("unrecognized argument: %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(mirrorCmd)
}
