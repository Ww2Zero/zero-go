package _8_isValidBST

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return validBSTHelper(root, nil, nil)
}

func validBSTHelper(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return validBSTHelper(root.Left, min, root) && validBSTHelper(root.Right, root, max)
}

//leetcode submit region end(Prohibit modification and deletion)
