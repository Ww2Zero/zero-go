package _020_numEnclaves

import (
	"testing"
)

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				// [0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]
				grid: [][]int{
					{0, 0, 0, 0},
					{1, 0, 1, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 0},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves(tt.args.grid); got != tt.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
