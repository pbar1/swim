package cli

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cobra"

	"github.com/pbar1/swim/pkg/model"
	"github.com/pbar1/swim/pkg/ncaa"
)

// getCmd represents the create command
var ncaamirrorCmd = &cobra.Command{
	Use:   "ncaamirror",
	Short: "Download a lot of times",
	Long:  `Download a lot of times`,
	Run: func(cmd *cobra.Command, args []string) {

		var wg sync.WaitGroup
		pool, err := ants.NewPool(12)
		if err != nil {
			log.Fatal(err)
		}
		defer pool.Release()

		for _, swimEvent := range model.AllSwimEvents {
			if swimEvent.Course == model.CourseSCY {

				for _, gender := range []model.Gender{model.GenderMale, model.GenderFemale} {
					for _, namedDateRange := range ncaa.NamedDateRangesD3 {
						for _, conference := range ncaa.ConferencesD3 {

							job := func() {
								params := ncaa.EventRankSearchParameters{
									Conference:     conference,
									NamedDateRange: namedDateRange,
									Distance:       swimEvent.Distance,
									Stroke:         swimEvent.Stroke,
									Course:         swimEvent.Course,
									Gender:         gender,
									Standard:       "NS",
									MaxResults:     "7000",
								}

								status := ncaaMirrorJobStatus{
									Conference:     conference,
									NamedDateRange: namedDateRange,
									Distance:       swimEvent.Distance,
									Stroke:         swimEvent.Stroke,
									Course:         swimEvent.Course,
									Gender:         gender,
									Standard:       "NS",
									MaxResults:     "7000",
								}

								results, err := ncaa.EventRankSearchD3(&params, 8*time.Second, true)
								if err != nil {
									status.Error = err
								}
								if results == nil {
									status.ReturnedResults = 0
								} else {
									status.ReturnedResults = len(results)
								}

								statusCSV, err := gocsv.MarshalString([]ncaaMirrorJobStatus{status})
								if err != nil {
									log.Printf("unable to marshal status to csv: %v", err)
									return
								} else {
									fmt.Println(statusCSV)
								}

								resultsCSV, err := gocsv.MarshalBytes(results)
								if err != nil {
									log.Printf("unable to marshal results to csv: %v", err)
									return
								}

								h := sha1.New()
								h.Write(resultsCSV)
								bs := h.Sum(nil)

								fname := fmt.Sprintf("results/%x.csv", bs)
								if err := ioutil.WriteFile(fname, resultsCSV, 0644); err != nil {
									log.Printf("unable to write results file: %v", err)
								}

								wg.Done()
							}

							if err := pool.Submit(job); err != nil {
								log.Printf("unable to submit job: %v", err)
							}
							wg.Add(1)

						}
					}
				}

			}
		}

		wg.Wait()
	},
}

type ncaaMirrorJobStatus struct {
	Conference      string       `csv:"Conference"`
	NamedDateRange  string       `csv:"NamedDateRange"`
	Distance        int          `csv:"Distance"`
	Stroke          model.Stroke `csv:"Stroke"`
	Course          model.Course `csv:"Course"`
	Gender          model.Gender `csv:"Gender"`
	Standard        string       `csv:"Standard"`
	MaxResults      string       `csv:"MaxResults"`
	ReturnedResults int          `csv:"ReturnedResults"`
	Error           error        `csv:"Error"`
}

func init() {
	rootCmd.AddCommand(ncaamirrorCmd)
}
