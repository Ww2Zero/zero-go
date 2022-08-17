package _302_deepestLeavesSum

//leetcode submit region begin(Prohibit modification and deletion)

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	sum := 0
	for len(q) > 0 {
		var p []*TreeNode
		sum = 0
		for i := 0; i < len(q); i++ {
			node := q[i]
			sum = sum + node.Val
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
