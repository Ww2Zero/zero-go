package _14_flatten

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	tmp := root
	for tmp.Right != nil {
		tmp = tmp.Right
	}
	tmp.Right = right

	return
}

//leetcode submit region end(Prohibit modification and deletion)
