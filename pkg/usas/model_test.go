package usas

import (
	"testing"
)

func TestGender_String(t *testing.T) {
	tests := []struct {
		name string
		g    Gender
		want string
	}{
		{"Stringify Gender enum", GenderMale, "Male"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.String(); got != tt.want {
				t.Errorf("Gender.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance_String(t *testing.T) {
	tests := []struct {
		name string
		d    Distance
		want string
	}{
		{"Stringify Distance enum", Distance50, "50"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Distance.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStroke_String(t *testing.T) {
	tests := []struct {
		name string
		s    Stroke
		want string
	}{
		{"Stringify Stroke enum", StrokeFly, "FL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Stroke.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourse_String(t *testing.T) {
	tests := []struct {
		name string
		c    Course
		want string
	}{
		{"Stringify Course enum", CourseSCY, "SCY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Course.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZone_String(t *testing.T) {
	tests := []struct {
		name string
		z    Zone
		want string
	}{
		{"Stringify Zone enum", ZoneAll, "All"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.z.String(); got != tt.want {
				t.Errorf("Zone.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStroke_ID(t *testing.T) {
	tests := []struct {
		name string
		s    Stroke
		want string
	}{
		{"Get ID string for Stroke enum", StrokeFly, "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ID(); got != tt.want {
				t.Errorf("Stroke.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourse_ID(t *testing.T) {
	tests := []struct {
		name string
		c    Course
		want string
	}{
		{"Get ID string for Course enum", CourseSCY, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ID(); got != tt.want {
				t.Errorf("Course.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZone_ID(t *testing.T) {
	tests := []struct {
		name string
		z    Zone
		want string
	}{
		{"Get ID string for Zone enum", ZoneAll, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.z.ID(); got != tt.want {
				t.Errorf("Zone.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
