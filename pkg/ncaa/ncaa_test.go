package ncaa

import (
	"reflect"
	"testing"
	"time"
)

func TestEventRankSearchD3(t *testing.T) {
	type args struct {
		params   *EventRankSearchParameters
		timeout  time.Duration
		headless bool
	}
	tests := []struct {
		name    string
		args    args
		want    []SwimResult
		wantErr bool
	}{
		{
			"CanFindMe",
			args{
				params: &EventRankSearchParameters{
					Conference:     "Heartland Collegiate",
					NamedDateRange: "(4) 2016-17 NCAA Div III",
					Distance:       "100",
					Stroke:         "FL",
					Course:         "SCY",
					Gender:         "Male",
					Standard:       "NS",
					MaxResults:     "1",
				},
				timeout:  15 * time.Second,
				headless: true,
			},
			[]SwimResult{
				{
					AltAdjSwimTime: "49.08",
					ConvFrom:       "",
					Name:           "Bartine, Pierce",
					Age:            21,
					School:         "Rose-Hulman",
					Year:           "Senior",
					MeetName:       "2017 CCIW Championships - NCAA",
					SwimDate:       "2017-02-10",
					Standard:       "B",
					Gender:         "Male",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EventRankSearchD3(tt.args.params, tt.args.timeout, tt.args.headless)
			if (err != nil) != tt.wantErr {
				t.Errorf("EventRankSearchD3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventRankSearchD3() got = %v, want %v", got, tt.want)
			}
		})
	}
}
