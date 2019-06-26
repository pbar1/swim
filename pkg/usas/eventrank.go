package usas

import (
	"compress/gzip"
	"net/url"

	"golang.org/x/net/html"
)

// EventRank performs an Event Rank / Top Times search
func EventRank() error {
	reqURL := "https://www.usaswimming.org/times/event-rank-search/ListEventRankResultsForFilter"
	formData := url.Values{
		"divId":                    {"Times_EventRankSearch_Index_Div_1"},
		"SelectedDateType":         {"DateRange"},
		"StartDate":                {"6/11/2019"},
		"EndDate":                  {"6/12/2019"},
		"DateRangeID":              {"0"},
		"SelectedGender":           {"Male"},
		"DSC[DistanceID]":          {"100"},
		"DSC[StrokeID]":            {"3"},
		"DSC[CourseID]":            {"3"},
		"StandardID":               {"12"}, // Slower than B
		"LSCs":                     {"'All'"},
		"ZoneID":                   {"0"}, // All
		"AgeRangeStart":            {"18"},
		"AgeRangeEnd":              {"18"},
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

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer reader.Close()

	tok := html.NewTokenizer(reader)
	for {
		tt := tok.Next()

	}

	return nil
}
