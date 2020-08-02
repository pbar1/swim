package ncaa

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"

	"github.com/pbar1/swim/pkg/model"
	"github.com/pbar1/swim/pkg/util"
)

func EventRankSearchD1(params *EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	return eventRankSearch(URLEventRankD1, model.OrgNCAADiv1, params, timeout, headless)
}

func EventRankSearchD2(params *EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	return eventRankSearch(URLEventRankD2, model.OrgNCAADiv2, params, timeout, headless)
}

func EventRankSearchD3(params *EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	return eventRankSearch(URLEventRankD3, model.OrgNCAADiv3, params, timeout, headless)
}

func eventRankSearch(url, org string, params *EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		func(a *chromedp.ExecAllocator) {
			chromedp.UserAgent(UserAgent)
			chromedp.Flag("headless", headless)(a)
		},
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	go func() {
		time.Sleep(timeout)
		cancel()
	}()

	var xPathGender string
	if params.Gender.String() == "Male" {
		xPathGender = XPathGenderMale
	} else if params.Gender.String() == "Female" {
		xPathGender = XPathGenderFemale
	} else {
		return nil, fmt.Errorf("unknown gender: %s", params.Gender)
	}

	var rawHTML string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(XPathConference),
		chromedp.SendKeys(XPathConference, params.Conference),
		chromedp.WaitVisible(XPathNamedDateRange),
		chromedp.SendKeys(XPathNamedDateRange, params.NamedDateRange),
		chromedp.WaitVisible(XPathEventIndDist),
		chromedp.SendKeys(XPathEventIndDist, fmt.Sprintf("%d", params.Distance)),
		chromedp.WaitVisible(XPathEventIndStroke),
		chromedp.SendKeys(XPathEventIndStroke, params.Stroke.String()),
		chromedp.WaitVisible(XPathEventIndCourse),
		chromedp.SendKeys(XPathEventIndCourse, params.Course.String()),
		chromedp.WaitVisible(xPathGender),
		chromedp.Click(xPathGender),
		chromedp.WaitVisible(XPathCut),
		chromedp.SendKeys(XPathCut, params.Standard),
		chromedp.WaitVisible(XPathIncludeAllTimes),
		chromedp.Click(XPathIncludeAllTimes),
		chromedp.WaitVisible(XPathMaxResults),
		chromedp.Clear(XPathMaxResults),
		chromedp.SendKeys(XPathMaxResults, params.MaxResults),
		chromedp.WaitVisible(XPathSearchButton),
		chromedp.Click(XPathSearchButton),
		chromedp.WaitVisible(XPathSearchResultsTitle),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			rawHTML, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	})

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHTML))
	if err != nil {
		return nil, err
	}

	tableHTML, err := goquery.OuterHtml(doc.Find(SelectorSearchResults).First())
	if err != nil {
		return nil, err
	}

	var table rawSwimResultTable
	if err := xml.Unmarshal([]byte(tableHTML), &table); err != nil {
		return nil, err
	}

	swimResults := make([]model.SwimResult, 0, 0)
	for _, r := range table.Results {
		if r.Class == "DataGridItemStyle" || r.Class == "DataGridAlternatingItemStyle" {

			// this is done because the NCAA search calculates and displays the swimmer's age as of today's date
			swimDate, err := time.Parse("1/2/2006", r.Field[9])
			if err != nil {
				return nil, err
			}
			todayAge, err := strconv.Atoi(r.Field[5])
			if err != nil {
				return nil, err
			}
			thenAge := todayAge - int(math.Ceil(time.Since(swimDate).Hours()/8760.0))

			altAdjTime := r.Field[2]
			recordedTime := strings.TrimSpace(r.Field[3])
			altAdjSecs, err := util.ParseSeconds(altAdjTime)
			if err != nil {
				return nil, err
			}
			var recordedSecs float64
			if recordedTime == "" {
				recordedTime = altAdjTime
				recordedSecs = altAdjSecs
			} else {
				recordedSecs, err = util.ParseSeconds(recordedTime)
				if err != nil {
					return nil, err
				}
			}

			swimResults = append(swimResults, model.SwimResult{
				SwimEvent:             &model.SwimEvent{Distance: params.Distance, Stroke: params.Stroke, Course: params.Course},
				SwimID:                -1,
				SwimTime:              recordedTime,
				SwimTimeSeconds:       recordedSecs,
				AltAdjSwimTime:        altAdjTime,
				AltAdjSwimTimeSeconds: altAdjSecs,
				PowerPoints:           -1, // TODO, write a util to fill this in
				Relay:                 false,
				Organization:          org,
				Locale:                params.Conference,
				TeamName:              r.Field[6],
				SwimmerName:           r.Field[4],
				SwimmerID:             -1,
				Gender:                params.Gender,
				Age:                   thenAge,
				MeetName:              r.Field[8],
				MeetID:                -1,
				SwimDate:              &model.SwimDate{T: swimDate},
				TimeStandard:          r.Field[10],
				USASForeign:           false,
				USASSponsorImage:      "",
				USASSponsorWebsite:    "",
				NCAAClassYear:         "",
				NCAAIneligible:        false, // TODO, actually detect this
			})
		}
	}

	return swimResults, nil
}
