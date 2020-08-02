package ncaa

import "github.com/pbar1/swim/pkg/model"

type (
	NamedDateRange string

	SwimResult struct {
		AltAdjSwimTime string // pos 2
		ConvFrom       string // pos 3
		Name           string // pos 4
		Age            int    // pos 5, make sure to correct returned age
		School         string // pos 6
		Year           string // pos 7
		MeetName       string // pos 8
		SwimDate       string // pos 9
		Standard       string // pos 10
		Gender         string
	}

	EventRankSearchParameters struct {
		Conference     string
		NamedDateRange string
		Distance       int
		Stroke         model.Stroke
		Course         model.Course
		Gender         model.Gender
		Standard       string
		MaxResults     string
	}

	rawSwimResultTable struct {
		Results []*rawSwimResult `xml:"tbody>tr"`
	}

	rawSwimResult struct {
		Class string   `xml:"class,attr"`
		Field []string `xml:"td"`
	}
)

const (
	UserAgent = `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`

	URLEventRankD1 = "https://legacy.usaswimming.org/DesktopDefault.aspx?TabId=2974"
	URLEventRankD2 = "https://legacy.usaswimming.org/DesktopDefault.aspx?TabId=3049"
	URLEventRankD3 = "https://legacy.usaswimming.org/DesktopDefault.aspx?TabId=3055"

	XPathConference         = `//*[@id="ctl02_ddLevel1"]`
	XPathNamedDateRange     = `//*[@id="ctl02_ddNamedDateRange"]`
	XPathEventIndDist       = `//*[@id="ctl02_ucDistanceStrokeCourseCtl_ddDistance"]`
	XPathEventIndStroke     = `//*[@id="ctl02_ucDistanceStrokeCourseCtl_ddStroke"]`
	XPathEventIndCourse     = `//*[@id="ctl02_ucDistanceStrokeCourseCtl_ddCourse"]`
	XPathGenderMale         = `//*[@id="ctl02_rbGenderMale"]`
	XPathGenderFemale       = `//*[@id="ctl02_rbGenderFemale"]`
	XPathCut                = `//*[@id="ctl02_ddStandard"]`
	XPathIncludeAllTimes    = `//*[@id="ctl02_radAllTimesForSwimmer"]`
	XPathMaxResults         = `//*[@id="ctl02_txtMaxResults"]`
	XPathSearchButton       = `//*[@id="ctl02_btnSearch"]`
	XPathSearchResultsTitle = `//*[@id="ctl02_trSearchResultsTitle"]`
	XPathNoTimesFound       = `//*[@id="ctl02_lblNoSearchResults"]`

	SelectorSearchResults = `#ctl02_dgSearchResults`
)

var (
	NamedDateRangesD3 = []string{
		"(1) 2019-20 NCAA Div III",
		"(2) 2018-19 NCAA Div III",
		"(3) 2017-18 NCAA Div III",
		"(4) 2016-17 NCAA Div III",
		"(5) 2015-16 NCAA Div III",
		"(6) 2014-15 NCAA Div III",
		"(7) 2013-14 NCAA Div III",
		"(8) 2012-13 NCAA Div III",
		"(9) 2011-12 NCAA Div III",
		"(10) 2010-11 NCAA Div III",
	}

	ConferencesD3 = []string{
		// "-- All --",
		"Allegheny Mountain",
		"American Southwest",
		"Appalachian (ASC)",
		"Atlantic East Conference",
		"Bluegrass Mountain",
		"Capital Athletic",
		"Centennial",
		"City Univ. of New York",
		"College of Illinois/Wisc",
		"Colonial States Athletic",
		"Commonwealth Coast",
		"Empire 8",
		"Great Northeast Athletic",
		"Great South Athletic",
		"Heartland Collegiate",
		"Independent",
		"Iowa Intercollegiate",
		"Landmark",
		"Liberal Arts",
		"Liberty League",
		"Little East",
		"Massachusetts State",
		"Metropolitan Swim",
		"Michigan Intercollegiate",
		"Middle Atlantic",
		"Midwest",
		"Minnesota Intercollegiate",
		"New England Intercoll.",
		"New England Small Coll",
		"New England Women's/Men's",
		"New Jersey Athletic",
		"North Atlantic",
		"North Coast Athletic",
		"North Eastern Athletic",
		"Northern Athletics",
		"Northwest Conference",
		"Ohio Athletic",
		"Old Dominion Athletic",
		"Pacific Collegiate",
		"Presidents' Athletic Conf",
		"Skyline",
		"Southern Athletic Associa",
		"Southern California",
		"Southern Collegiate",
		"St. Louis Intercollegiate",
		"State Univ of New York",
		"University Athletic",
		"Upper Midwest Athletic",
		"USA South Athletic",
		"Wisconsin Intercollegiate",
	}
)
