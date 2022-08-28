package _01_insertIntoBST

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	node := root
	var parentNode *TreeNode
	for node != nil {
		parentNode = node
		if val > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	if val > parentNode.Val {
		parentNode.Right = &TreeNode{Val: val}
	} else {
		parentNode.Left = &TreeNode{Val: val}
	}
	return root
}
