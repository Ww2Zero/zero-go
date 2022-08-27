package _04_maxDepth

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	deep := 0

	if root == nil {
		return deep
	}

	var q []*TreeNode

	q = append(q, root)

	for len(q) > 0 {

		var p []*TreeNode

		deep++
		for i := 0; i < len(q); i++ {
			node := q[i]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return deep

}
