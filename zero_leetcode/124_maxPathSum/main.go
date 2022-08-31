package _24_maxPathSum

import "math"

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt

	var maxPathSumHelper func(node *TreeNode) int
	maxPathSumHelper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := max(maxPathSumHelper(node.Left), 0)
		rightSum := max(maxPathSumHelper(node.Right), 0)

		res = max(node.Val+leftSum+rightSum, res)

		return node.Val + max(leftSum, rightSum)
	}

	maxPathSumHelper(root)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
