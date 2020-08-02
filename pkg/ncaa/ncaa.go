package ncaa

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"

	"github.com/pbar1/swim/pkg/model"
	"github.com/pbar1/swim/pkg/util"
)

var (
	recordedTimeOpenReg  *regexp.Regexp
	recordedTimeCloseReg *regexp.Regexp
)

func init() {
	var err error
	recordedTimeOpenReg, err = regexp.Compile(`<td style="white-space:nowrap;">\s*<span id="ctl02_dgSearchResults_lblSwimTimeFormatted_[0-9]+">`)
	if err != nil {
		log.Fatal(err)
	}
	recordedTimeCloseReg, err = regexp.Compile(`</span>\n                            </td>`)
	if err != nil {
		log.Fatal(err)
	}
}

func EventRankSearchD1(params EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	return eventRankSearch(URLEventRankD1, model.OrgNCAADiv1, params, timeout, headless)
}

func EventRankSearchD2(params EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	return eventRankSearch(URLEventRankD2, model.OrgNCAADiv2, params, timeout, headless)
}

func EventRankSearchD3(params EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
	return eventRankSearch(URLEventRankD3, model.OrgNCAADiv3, params, timeout, headless)
}

func eventRankSearch(url, org string, params EventRankSearchParameters, timeout time.Duration, headless bool) ([]model.SwimResult, error) {
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
	if params.Gender == model.GenderMale {
		xPathGender = XPathGenderMale
	} else if params.Gender == model.GenderFemale {
		xPathGender = XPathGenderFemale
	} else {
		return nil, fmt.Errorf("unknown gender: %s", params.Gender)
	}

	// hack for sendkeys behavior of 50 and 500
	sendKeysDistance := ""
	if params.Distance != 50 {
		sendKeysDistance = fmt.Sprintf("%d", params.Distance)
	}

	var rawHTML string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(XPathConference),
		chromedp.SendKeys(XPathConference, params.Conference),
		chromedp.WaitVisible(XPathNamedDateRange),
		chromedp.SendKeys(XPathNamedDateRange, params.NamedDateRange),
		chromedp.WaitVisible(XPathEventIndDist),
		chromedp.SendKeys(XPathEventIndDist, sendKeysDistance),
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
				return fmt.Errorf("getting dom: %v", err)
			}
			rawHTML, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return fmt.Errorf("getting response html: %v", err)
		}),
	})

	swimResults := make([]model.SwimResult, 0, 0)
	if strings.Contains(rawHTML, TextNoTimesFound) {
		return swimResults, nil
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHTML))
	if err != nil {
		return nil, fmt.Errorf("making goquery document: %v", err)
	}

	tableHTML, err := goquery.OuterHtml(doc.Find(SelectorSearchResults).First())
	if err != nil {
		return nil, fmt.Errorf("getting results table html: %v", err)
	}

	// hack to clean xml for parsing recorded time and eligibility
	tableHTML = recordedTimeOpenReg.ReplaceAllString(tableHTML, "<td>")
	tableHTML = recordedTimeCloseReg.ReplaceAllString(tableHTML, "</td>")

	var table rawSwimResultTable
	if err := xml.Unmarshal([]byte(tableHTML), &table); err != nil {
		return nil, fmt.Errorf("unmarshal xml results table: %v", err)
	}

	for _, r := range table.Results {
		if r.Class == "DataGridItemStyle" || r.Class == "DataGridAlternatingItemStyle" {

			// this is done because the NCAA search calculates and displays the swimmer's age as of today's date
			swimDate, err := time.Parse("1/2/2006", r.Field[9])
			if err != nil {
				return nil, fmt.Errorf("parse swim date '%s': %v", r.Field[9], err)
			}
			todayAge, err := strconv.Atoi(r.Field[5])
			if err != nil {
				return nil, fmt.Errorf("convert todayAge to int %s: %v", r.Field[5], err)
			}
			thenAge := todayAge - int(math.Ceil(time.Since(swimDate).Hours()/8760.0))

			ineligible := false
			recordedTime := r.Field[1]
			if strings.Contains(recordedTime, "*") {
				ineligible = true
				recordedTime = strings.Replace(recordedTime, " *", "", 1)
			}
			recordedSecs, err := util.ParseSeconds(recordedTime)
			if err != nil {
				return nil, fmt.Errorf("parse seconds %s: %v", recordedTime, err)
			}

			altAdjTime := r.Field[2]
			altAdjSecs, err := util.ParseSeconds(altAdjTime)
			if err != nil {
				return nil, fmt.Errorf("parse seconds %s: %v", altAdjTime, err)
			}

			// TODO: ConvFrom is actually of the form ex. "SCM"
			swimResults = append(swimResults, model.SwimResult{
				SwimEvent:             model.SwimEvent{Distance: params.Distance, Stroke: params.Stroke, Course: params.Course},
				SwimID:                -1,
				SwimTime:              recordedTime,
				SwimTimeSeconds:       recordedSecs,
				AltAdjSwimTime:        altAdjTime,
				AltAdjSwimTimeSeconds: altAdjSecs,
				PowerPoints:           -1, // TODO: write a util to fill this in
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
				SwimDate:              model.SwimDate{T: swimDate},
				TimeStandard:          r.Field[10],
				USASForeign:           false,
				USASSponsorImage:      "",
				USASSponsorWebsite:    "",
				NCAAClassYear:         r.Field[7],
				NCAAIneligible:        ineligible,
				NCAAConvFrom:          strings.TrimSpace(r.Field[3]),
			})
		}
	}

	return swimResults, nil
}
