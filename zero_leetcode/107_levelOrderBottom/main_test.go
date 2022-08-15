package _07_levelOrderBottom

import (
	"math"
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test",
			args: args{root: newTree([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})},
			want: [][]int{[]int{15, 7}, []int{9, 20}, []int{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTree(ints []int) *TreeNode {
	if len(ints) == 0 {
		return nil
	}
	return build(ints, 0)
}

func build(ints []int, n int) *TreeNode {
	if len(ints) == 0 {
		return nil
	}
	if n >= len(ints) {
		return nil
	}
	if ints[n] == math.MinInt {
		return nil
	}
	node := &TreeNode{Val: ints[n]}
	node.Left = build(ints, 2*n+1)
	node.Right = build(ints, n*2+2)
	return node
}
