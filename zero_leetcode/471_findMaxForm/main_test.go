package _71_findMaxForm

import "testing"

//[[0 0 0 0]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 1 1]]

//[[0 0 0 0]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 2 2]
// [0 1 2 2]]

//[[0 0 0 0]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 1 1]
// [0 1 2 2]
// [0 1 2 2]]

//[[0 1 1 1]
// [0 1 2 2]
// [0 1 2 2]
// [0 1 2 2]
// [0 1 2 3]
// [0 1 2 3]]

//[[0 1 1 1]
// [1 2 2 2]
// [1 2 3 3]
// [1 2 3 3]
// [1 2 3 3]
// [1 2 3 4]]
func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
