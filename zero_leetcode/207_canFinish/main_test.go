package _07_canFinish

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{[]int{1, 0}},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				numCourses: 6,
				// [[3, 0], [3, 1], [4, 1], [4, 2], [5, 3], [5, 4]]
				prerequisites: [][]int{{3, 0}, {3, 1}, {4, 1}, {4, 2}, {5, 3}, {5, 4}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
