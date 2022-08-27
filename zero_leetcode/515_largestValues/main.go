package _15_largestValues

import "math"

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		var p []*TreeNode
		var sub []int
		for i := 0; i < len(q); i++ {
			node := q[i]
			sub = append(sub, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
		res = append(res, max(sub...))
	}

	return
}

func max(nums ...int) int {
	m := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m

}
