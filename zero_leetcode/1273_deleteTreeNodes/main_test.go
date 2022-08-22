package _273_deleteTreeNodes

import "testing"

func Test_deleteTreeNodes(t *testing.T) {
	type args struct {
		nodes  int
		parent []int
		value  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nodes:  7,
				parent: []int{-1, 0, 0, 1, 2, 2, 2},
				value:  []int{1, -2, 4, 0, -2, -1, -1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteTreeNodes(tt.args.nodes, tt.args.parent, tt.args.value); got != tt.want {
				t.Errorf("deleteTreeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
