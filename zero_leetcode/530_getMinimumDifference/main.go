package _30_getMinimumDifference

import "math"

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// getMinimumDifference 513. 找到二叉搜索树中的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	minVal := math.MaxInt64

	var travel func(node *TreeNode)

	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < minVal {
			minVal = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)

	}
	travel(root)
	return minVal
}
