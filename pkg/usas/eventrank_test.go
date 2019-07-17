package usas

import (
	"testing"
	"time"
)

func TestEventRank(t *testing.T) {
	type args struct {
		in *EventRankInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Event Rank Search returns data",
			args: args{in: &EventRankInput{
				StartDate: time.Date(2007, time.September, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2008, time.August, 31, 0, 0, 0, 0, time.UTC),
				Gender:    GenderMale,
				Distance:  Distance100,
				Stroke:    StrokeFly,
				Course:    CourseLCM,
				Zone:      ZoneAll,
				StartAge:  18,
				EndAge:    40,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EventRank(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("EventRank() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
