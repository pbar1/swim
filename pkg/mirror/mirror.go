package mirror

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/panjf2000/ants/v2"

	"github.com/pbar1/swim/pkg/model"
	"github.com/pbar1/swim/pkg/ncaa"
	"github.com/pbar1/swim/pkg/util"
)

func MirrorNCAAD1(poolSize int, timeout time.Duration) {
	mirrorNCAA(model.OrgNCAADiv1, ncaa.NamedDateRangesD1, ncaa.ConferencesD1, poolSize, timeout)
}

func MirrorNCAAD2(poolSize int, timeout time.Duration) {
	mirrorNCAA(model.OrgNCAADiv2, ncaa.NamedDateRangesD2, ncaa.ConferencesD2, poolSize, timeout)
}

func MirrorNCAAD3(poolSize int, timeout time.Duration) {
	mirrorNCAA(model.OrgNCAADiv3, ncaa.NamedDateRangesD3, ncaa.ConferencesD3, poolSize, timeout)
}

func mirrorNCAA(div string, namedDateRanges, conferences []string, poolSize int, timeout time.Duration) {
	// setup job pool
	var wg sync.WaitGroup
	pool, err := ants.NewPool(poolSize)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Release()

	// choose search function for proper division
	var searchFn func(ncaa.EventRankSearchParameters, time.Duration, bool) ([]model.SwimResult, error)
	switch div {
	case model.OrgNCAADiv1:
		searchFn = ncaa.EventRankSearchD1
	case model.OrgNCAADiv2:
		searchFn = ncaa.EventRankSearchD2
	case model.OrgNCAADiv3:
		searchFn = ncaa.EventRankSearchD3
	default:
		log.Fatalf("unknown ncaa division: %s", div)
	}

	// enumerate total search queryspace for SCY
	for _, swimEvent := range model.AllSwimEvents {
		if swimEvent.Course == model.CourseSCY {
			for _, gender := range []model.Gender{model.GenderMale, model.GenderFemale} {
				for _, namedDateRange := range namedDateRanges {
					for _, conference := range conferences {

						job := func() {
							defer wg.Done()
							conference := conference
							namedDateRange := namedDateRange
							gender := gender
							swimEvent := swimEvent
							cnf := strings.ToLower(util.RemoveSpecialChars(conference))
							nd := strings.Split(namedDateRange, " ")[1]
							ev := strings.ReplaceAll(swimEvent.String(), " ", "-")
							qid := strings.ToLower(fmt.Sprintf("%s-%s-%s-%s-%s", div, gender, ev, nd, cnf))

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

							results, err := searchFn(params, timeout, true)
							if err != nil {
								log.Printf("event rank search error: qid=%s, err=%v", qid, err)
								return
							}
							if results == nil {
								results = make([]model.SwimResult, 0, 0)
							}
							log.Printf("qid=%s, numResults=%d, err=%v", qid, len(results), err)

							resultsCSV, err := gocsv.MarshalBytes(results)
							if err != nil {
								log.Printf("csv marshal results error: qid=%s, err=%v", qid, err)
								return
							}

							fname := fmt.Sprintf("results/%s.csv", qid)
							if err := ioutil.WriteFile(fname, resultsCSV, 0644); err != nil {
								log.Printf("results file write error: qid=%s, err=%v", qid, err)
							}
						}

						if err := pool.Submit(job); err != nil {
							log.Printf("submit job to pool error: %v", err)
						}
						wg.Add(1)

					} // end for conference
				} // end for namedDateRange
			} // end for gender
		} // end if SCY
	} // end for swimEvent

	wg.Wait()
}
