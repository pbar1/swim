package usas

import (
	"compress/gzip"
	"io"
	"net/url"
	"os"
	"time"
)

type EventRankInput struct {
	StartDate time.Time
	EndDate   time.Time
	Gender    Gender
	Distance  Distance
	Stroke    Stroke
	Course    Course
	Zone      Zone
	StartAge  Age
	EndAge    Age
}

// EventRank performs an Event Rank / Top Times search
func EventRank(in *EventRankInput) error {
	reqURL := "https://www.usaswimming.org/times/event-rank-search/ListEventRankResultsForFilter"

	formData := url.Values{
		"divId":                    {"Times_EventRankSearch_Index_Div_1"},
		"SelectedDateType":         {"DateRange"},
		"StartDate":                {in.StartDate.Format("1/2/2006")},
		"EndDate":                  {in.EndDate.Format("1/2/2006")},
		"DateRangeID":              {"0"},
		"SelectedGender":           {string(in.Gender)},
		"DSC[DistanceID]":          {string(in.Distance)},
		"DSC[StrokeID]":            {in.Stroke.ID()},
		"DSC[CourseID]":            {in.Course.ID()},
		"StandardID":               {"12"}, // Slower than B
		"LSCs":                     {"'All'"},
		"ZoneID":                   {in.Zone.ID()}, // All
		"AgeRangeStart":            {string(in.StartAge)},
		"AgeRangeEnd":              {string(in.EndAge)},
		"SelectedTimesToInclude":   {"All"}, // Other option is "Best"
		"SelectedMembersToInclude": {"No"},  // Include non USA Swimming members
		"MaxResults":               {"5000"},
		"OrderBy":                  {"Rank"},
		"clubId":                   {""},
		"TimeType":                 {"Individual"},
	}

	resp, err := makePost(reqURL, formData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer reader.Close()

	// tok := html.NewTokenizer(reader)
	// for {
	// 	tt := tok.Next()
	// }

	io.Copy(os.Stdout, reader)

	return nil
}
