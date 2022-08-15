package _03_zigzagLevelOrder

import (
	"math"
	"reflect"
	"testing"
)

//解答失败: 测试用例:[1,2,3,4,null,null,5] 测试结果:[[1],[3,2],[5,4]] 期望结果:[[1],[3,2],[4,5]] stdout:
func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{root: newTree([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})},
			want: [][]int{[]int{3}, []int{20, 9}, []int{15, 7}},
		},
		{
			name: "test2",
			args: args{root: newTree([]int{1, 2, 3, 4, math.MinInt, math.MinInt, 5})},
			want: [][]int{[]int{1}, []int{3, 2}, []int{4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 0,1,2,3,4,5,6
// 1,2,3,4,5,6,7
// root  1
//    2    3
//  4   5 6  7

// idx 2-->5  2*1+2 = 4
// idx 3-->7  2*2+2 = 6

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
