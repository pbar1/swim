package cli

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/pbar1/swim/pkg/mirror"
)

var (
	pool    int
	timeout time.Duration
)

// getCmd represents the create command
var mirrorCmd = &cobra.Command{
	Use:   "mirror [DIVISION]",
	Short: "Download all times available for a given parameter space",
	Long: `Enumerates all possible NCAA Event Rank Search parameters for SCY. This
includes both male and female, all named date ranges, all conferences, and all legal
competition events. Results are saved to disk in CSV files corresponding to the
search parameters that were used for that file. No results file will be written if
the query returns an error.

Values for DIVISION are one of: d1|d2|d3`,
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
	mirrorCmd.Flags().IntVarP(&pool, "pool-size", "p", 10, "Number of concurrent threads")
	mirrorCmd.Flags().DurationVarP(&timeout, "search-timeout", "t", 6*time.Second, "Seconds to allow searches before timing out")
}
