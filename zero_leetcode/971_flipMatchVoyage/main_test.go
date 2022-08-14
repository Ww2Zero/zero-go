package _71_flipMatchVoyage

import (
	"reflect"
	"testing"
)

func Test_flipMatchVoyage(t *testing.T) {
	type args struct {
		root   *TreeNode
		voyage []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				root:   newTree([]int{1, 2, 3}),
				voyage: []int{1, 3, 2},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipMatchVoyage(tt.args.root, tt.args.voyage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipMatchVoyage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTree(ints []int) *TreeNode {
	root := &TreeNode{Val: ints[0]}
	root.Left = &TreeNode{Val: ints[1]}
	root.Right = &TreeNode{Val: ints[2]}
	return root
}
