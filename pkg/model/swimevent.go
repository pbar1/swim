package model

import "fmt"

type SwimEvent struct {
	Distance int
	Stroke   Stroke
	Course   Course
}

func (e SwimEvent) String() string {
	return fmt.Sprintf("%s %s %d", e.Course, e.Stroke, e.Distance)
}

func (e SwimEvent) EventDesc() string {
	return fmt.Sprintf("%d %s %s", e.Distance, e.Stroke, e.Course)
}

func (e SwimEvent) IsValid() bool {
	for _, event := range AllSwimEvents {
		if e.Distance == event.Distance && e.Stroke == event.Stroke && e.Course == event.Course {
			return true
		}
	}
	return false
}

var (
	AllSwimEvents = []SwimEvent{
		// valid SCY events
		{50, StrokeFree, CourseSCY},
		{100, StrokeFree, CourseSCY},
		{200, StrokeFree, CourseSCY},
		{500, StrokeFree, CourseSCY},
		{1000, StrokeFree, CourseSCY},
		{1650, StrokeFree, CourseSCY},
		{50, StrokeBack, CourseSCY},
		{100, StrokeBack, CourseSCY},
		{200, StrokeBack, CourseSCY},
		{50, StrokeBreast, CourseSCY},
		{100, StrokeBreast, CourseSCY},
		{200, StrokeBreast, CourseSCY},
		{50, StrokeFly, CourseSCY},
		{100, StrokeFly, CourseSCY},
		{200, StrokeFly, CourseSCY},
		{100, StrokeIM, CourseSCY},
		{200, StrokeIM, CourseSCY},
		{400, StrokeIM, CourseSCY},

		// valid SCM events
		{50, StrokeFree, CourseSCM},
		{100, StrokeFree, CourseSCM},
		{200, StrokeFree, CourseSCM},
		{400, StrokeFree, CourseSCM},
		{800, StrokeFree, CourseSCM},
		{1500, StrokeFree, CourseSCM},
		{50, StrokeBack, CourseSCM},
		{100, StrokeBack, CourseSCM},
		{200, StrokeBack, CourseSCM},
		{50, StrokeBreast, CourseSCM},
		{100, StrokeBreast, CourseSCM},
		{200, StrokeBreast, CourseSCM},
		{50, StrokeFly, CourseSCM},
		{100, StrokeFly, CourseSCM},
		{200, StrokeFly, CourseSCM},
		{100, StrokeIM, CourseSCM},
		{200, StrokeIM, CourseSCM},
		{400, StrokeIM, CourseSCM},

		// valid LCM events
		{50, StrokeFree, CourseLCM},
		{100, StrokeFree, CourseLCM},
		{200, StrokeFree, CourseLCM},
		{400, StrokeFree, CourseLCM},
		{800, StrokeFree, CourseLCM},
		{1500, StrokeFree, CourseLCM},
		{50, StrokeBack, CourseLCM},
		{100, StrokeBack, CourseLCM},
		{200, StrokeBack, CourseLCM},
		{50, StrokeBreast, CourseLCM},
		{100, StrokeBreast, CourseLCM},
		{200, StrokeBreast, CourseLCM},
		{50, StrokeFly, CourseLCM},
		{100, StrokeFly, CourseLCM},
		{200, StrokeFly, CourseLCM},
		{200, StrokeIM, CourseLCM},
		{400, StrokeIM, CourseLCM},
	}
)
