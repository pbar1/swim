package usas

import "fmt"

type (
	// Gender is the gender of the swimmer
	Gender string

	// Distance is the length of the event
	Distance int

	// Stroke is the stroke of the event
	Stroke int

	// Course is the pool configuration of the event
	Course int

	// Age is the age of the swimmer
	Age int

	// Zone is the USA Swimming Zone ID
	Zone int
)

// List of constant values to use as enums
const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"

	Distance50   Distance = 50
	Distance100  Distance = 100
	Distance200  Distance = 200
	Distance400  Distance = 400
	Distance500  Distance = 500
	Distance800  Distance = 800
	Distance1000 Distance = 1000
	Distance1500 Distance = 1500
	Distance1650 Distance = 1650

	StrokeFree   Stroke = 1
	StrokeBack   Stroke = 2
	StrokeBreast Stroke = 3
	StrokeFly    Stroke = 4
	StrokeIM     Stroke = 5

	CourseSCY Course = 1
	CourseSCM Course = 2
	CourseLCM Course = 3

	ZoneAll      Zone = 0
	ZoneCentral  Zone = 1
	ZoneEastern  Zone = 2
	ZoneSouthern Zone = 3
	ZoneWestern  Zone = 4
)

var (
	strokes = [...]string{"FR", "BK", "BR", "FL", "IM"}
	courses = [...]string{"SCY", "SCM", "LCM"}
	zones   = [...]string{"All", "Central", "Eastern", "Southern", "Western"}
)

func (g Gender) String() string {
	return string(g)
}

func (d Distance) String() string {
	return fmt.Sprintf("%d", d)
}

func (s Stroke) String() string {
	return strokes[s-1]
}

func (c Course) String() string {
	return courses[c-1]
}

func (z Zone) String() string {
	return zones[z]
}

// ID returns the stroke ID
func (s Stroke) ID() string {
	return fmt.Sprintf("%d", s)
}

// ID returns the course ID
func (c Course) ID() string {
	return fmt.Sprintf("%d", c)
}

// ID returns the zone ID
func (z Zone) ID() string {
	return fmt.Sprintf("%d", z)
}
