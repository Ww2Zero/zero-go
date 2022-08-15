package _07_levelOrderBottom

//leetcode submit region begin(Prohibit modification and deletion)

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		var r [][]int
		var l []int
		var p []*TreeNode

		for i := 0; i < len(q); i++ {
			node := q[i]
			l = append(l, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		r = append(r, l)
		r = append(r, res...)
		res = r
		q = p
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
