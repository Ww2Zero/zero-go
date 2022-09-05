package _44_preorderTraversal

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				root: newTree([]int{1, -1, 2, 3}),
			},
			wantRes: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := preorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("preorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func newTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(nums); i += 2 {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != -1 {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		if i+1 < len(nums) && nums[i+1] != -1 {
			node.Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, node.Right)
		}
	}
	return root
}

func Test_preorderTraversal2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "test",
			args: args{
				root: newTree([]int{1, -1, 2, 3}),
			},
			wantRes: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.wantRes) {
				t.Errorf("preorderTraversal2() = %v, want %v", got, tt.wantRes)
			}
		})
	}
}
