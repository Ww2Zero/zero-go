package _38_convertBST

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// convertBST convert a BST to a greater sum tree
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
// 使每个节点 node 的新值等于原树中大于或等于node.val的值之和。
// 既然累加和是计算大于等于当前值的所有元素之和，那么每个节点都去计算右子树的和
func convertBST(root *TreeNode) *TreeNode {
	cnt := 0
	var sumTree func(node *TreeNode)
	sumTree = func(node *TreeNode) {
		if node == nil {
			return
		}
		sumTree(node.Right)
		cnt += node.Val
		node.Val = cnt
		sumTree(node.Left)
	}
	sumTree(root)
	return root
}
