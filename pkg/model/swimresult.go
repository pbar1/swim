package model

type SwimResult struct {
	// Competition event, ex. "LCM IM 400"
	SwimEvent *SwimEvent `csv:"SwimEvent"`

	// (Optional) Unique result ID
	SwimID int `csv:"SwimID"`

	// Recorded time, ex. "4:03.84"
	SwimTime string `csv:"SwimTime"`

	// Recorded time, ex. 243.84
	SwimTimeSeconds float64 `csv:"SwimTimeSeconds"`

	// Altitude-adjusted time, ex. "4:03.84"
	AltAdjSwimTime string `csv:"AltAdjSwimTime"`

	// Altitude-adjusted time, ex. 243.84
	AltAdjSwimTimeSeconds float64 `csv:"AltAdjSwimTimeSeconds"`

	// HY-TEK Points value, ex. 1120
	// more info: https://hytek.active.com/user_guides_html/swmm7/hy-tekpoints.htm
	PowerPoints int `csv:"PowerPoints"`

	// Whether swim was a relay lead-off
	Relay bool `csv:"Relay"`

	// Organization, one of: USAS|NCAA-D1|NCAA-D2|NCAA-D3
	Organization string `csv:"Organization"`

	// (Optional) LSC/conference for swimmer, ex. "MI"
	Locale string `csv:"Locale"`

	// Team/school for swimmer, ex. "Unattached"
	TeamName string `csv:"TeamName"`

	// Swimmer name, ex. "Phelps, Michael"
	SwimmerName string `csv:"SwimmerName"`

	// (Optional) Swimmer ID (if available), ex. 1034283
	SwimmerID int `csv:"SwimmerID"`

	// Swimmer gender, one of: Male|Female
	Gender Gender `csv:"Gender"`

	// Swimmer age on swim date, ex. 23
	Age int `csv:"Age"`

	// Swim meet name, ex. "2008 Olympic Games"
	MeetName string `csv:"MeetName"`

	// (Optional) Swim meet ID (if available), ex. 38805
	MeetID int `csv:"MeetID"`

	// Swim date, ex. "2008-08-09"
	SwimDate *SwimDate `csv:"SwimDate"`

	// (Optional) Org-specific time standard/cut achieved, ex. "Summer Nationals (LCM)"
	TimeStandard string `csv:"TimeStandard"`

	// (USA Swimming only) Whether swimmer is foreign to USA Swimming
	USASForeign bool `csv:"USASForeign"`

	// (USA Swimming only) Link to sponsor image
	USASSponsorImage string `csv:"USASSponsorImage"`

	// (USA Swimming only) Link to sponsor website
	USASSponsorWebsite string `csv:"USASSponsorWebsite"`

	// (NCAA only) Swimmer class/year, one of: Freshman|Sophomore|Junior|Senior
	NCAAClassYear string `csv:"NCAAClassYear"`

	// (NCAA) Whether the swim is ineligible for NCAA Championships
	NCAAIneligible bool `csv:"NCAAIneligible"`
}
