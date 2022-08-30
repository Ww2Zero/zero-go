package _32_findItinerary

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{
			name:    "test1",
			args:    args{tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}},
			wantRes: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findItinerary(tt.args.tickets); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findItinerary() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
