package model

import (
	"fmt"
	"time"
)

type (
	Gender       string
	Stroke       int
	Course       int
	Organization string

	SwimDate struct {
		T time.Time `csv:"-"`
	}
)

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"

	StrokeFree        Stroke = 1
	StrokeBack        Stroke = 2
	StrokeBreast      Stroke = 3
	StrokeFly         Stroke = 4
	StrokeIM          Stroke = 5
	StrokeFreeRelay   Stroke = 6
	StrokeMedleyRelay Stroke = 7

	CourseSCY Course = 1
	CourseSCM Course = 2
	CourseLCM Course = 3

	OrgUSASwimming = "USAS"
	OrgNCAADiv1    = "NCAA-D1"
	OrgNCAADiv2    = "NCAA-D2"
	OrgNCAADiv3    = "NCAA-D3"
)

var (
	strokes = []string{"FR", "BK", "BR", "FL", "IM", "FR-R", "MED-R"}
	courses = []string{"SCY", "SCM", "LCM"}
)

func (g Gender) String() string {
	return string(g)
}

func (s Stroke) String() string {
	return strokes[s-1]
}

func (c Course) String() string {
	return courses[c-1]
}

func (o Organization) String() string {
	return string(o)
}

func (d SwimDate) String() string {
	return d.T.Format("2006-01-02")
}

// ID returns the Stroke numerical ID as a string
func (s Stroke) ID() string {
	return fmt.Sprintf("%d", s)
}

// ID returns the Course numerical ID as a string
func (c Course) ID() string {
	return fmt.Sprintf("%d", c)
}
