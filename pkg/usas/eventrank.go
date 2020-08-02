package usas

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"time"
)

const (
	EventRankURL = "https://www.usaswimming.org/api/Times_TimesSearchTopTimesEventRankSearch/ListTimes"
)

type (
	EventRankInput struct {
		StartDate     time.Time
		EndDate       time.Time
		Gender        Gender
		Distance      Distance
		Stroke        Stroke
		Course        Course
		Zone          Zone
		StartAge      Age
		EndAge        Age
		BestTimesOnly bool
		MembersOnly   bool
		MaxResults    int
	}

	eventRankRaw struct {
		EventDesc         string
		EventSortOrder    int
		Rank              int
		SwimTime          string
		AltAdjSwimTime    string
		PowerPoints       int
		EventID           int
		FullName          string
		PersonClusteredID string
		Foreign           string
		Age               int
		LSC               string
		TeamName          string
		Athletes          interface{}
		AthletesPdf       interface{}
		MeetID            int
		MeetName          string
		SwimDate          string
		StandardName      string
		SponsorImage      string
		SponsorWebsite    string
		TimeId            int
		Distance          int
	}
)

// EventRank performs an Event Rank / Top Times search
func EventRank(in *EventRankInput) error {
	reqURL := EventRankURL

	lscs := "All"
	// if in.LSCs != nil && len(in.LSCs) > 0 {
	// 	var b strings.Builder
	// 	for _, lsc := range in.LSCs {
	// 		fmt.Fprintf(&b, "'%s',+", lsc)
	// 	}
	// 	lscs = strings.TrimSuffix(b.String(), ",+")
	// }

	bestTimesOnly := "All"
	if in.BestTimesOnly {
		bestTimesOnly = "Best"
	}

	membersOnly := "No"
	if in.MembersOnly {
		membersOnly = "Yes"
	}

	maxResults := "7000"
	if in.MaxResults > 0 && in.MaxResults < 7000 {
		maxResults = fmt.Sprintf("%d", in.MaxResults)
	}

	formData := url.Values{
		"DivId":                                 {"Times_TimesSearchTopTimesEventRankSearch_Index_Div-1"},
		"DateRangeId":                           {"0"},                             // must be 0 to enable FromDate and ToDate
		"FromDate":                              {in.StartDate.Format("1/2/2006")}, // must also specify EndDate
		"ToDate":                                {in.EndDate.Format("1/2/2006")},   // must also specify StartDate
		"TimeType":                              {"Individual"},                    // {Individual|Relay}, for filtering to relay lead swims? seemingly has no effect
		"DistanceId":                            {in.Distance.String()},
		"StrokeId":                              {in.Stroke.ID()},
		"CourseId":                              {in.Course.ID()},
		"StartAge":                              {fmt.Sprintf("%d", in.StartAge)}, // {All|1-50}, must also specify EndAge
		"EndAge":                                {fmt.Sprintf("%d", in.EndAge)},   // {All|1-50}, must also specify StartAge
		"Gender":                                {string(in.Gender)},
		"Standard":                              {"12"},        // 12 == "Slower than B" time standard cutoff, hardcoded for simplicity
		"IncludeTimesForUsaSwimmingMembersOnly": {membersOnly}, // {Yes|No}, where Yes returns only USA Swimming members without foreign swimmers
		"ClubId":                                {"-1"},        // uint club ID
		"ClubName":                              {""},          // string club name, like "Katy+Aquatic+Team+For+Youth"
		"Lscs":                                  {lscs},        // {...all the LSCs}, either "All" or "GU,PC" for example
		"Zone":                                  {in.Zone.ID()},
		"TimesToInclude":                        {bestTimesOnly}, // {All|Best}, where Best returns only top time per swimmer
		"SortBy1":                               {"EventSortOrder"},
		"SortBy2":                               {""},
		"SortBy3":                               {""},
		"MaxResults":                            {maxResults}, // must be less than 7000, though buggy returning greater then 5000 results
	}

	// perform http post
	resp, err := makePost(reqURL, formData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// unzip compressed response
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer reader.Close()

	// convert html stream into bytes
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return fmt.Errorf("unable to read into byte buffer: %v", err)
	}
	rawHtml := buf.Bytes()

	// parse html for data
	matcher := regexp.MustCompile(`(\bdata\b: )(\[.*])`)
	data := matcher.Find(rawHtml)
	if data == nil {
		return fmt.Errorf("no matches for data in html")
	}
	data = bytes.TrimPrefix(data, []byte("data: "))

	var dataParsed []eventRankRaw
	_ = json.Unmarshal(data, &dataParsed)
	fmt.Println(dataParsed)

	// TODO
	return nil
}
