package cli

import (
	"github.com/spf13/cobra"
)

// getCmd represents the create command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource",
	Long:  `Get a resource`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
