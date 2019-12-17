package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pbar1/swim/pkg/usas"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "swim",
	Short: "Swimming times utility",
	Long:  `Swimming times utilitsy`,
	Run: func(cmd *cobra.Command, args []string) {
		err := usas.EventRank(&usas.EventRankInput{
			StartDate: time.Now(),
			EndDate:   time.Now(),
			Gender:    usas.GenderMale,
			Distance:  usas.Distance100,
			Stroke:    usas.StrokeFree,
			Course:    usas.CourseSCY,
			Zone:      usas.ZoneAll,
			StartAge:  18,
			EndAge:    18,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv()
}
