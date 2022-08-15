package _44_preorderTraversal

//leetcode submit region begin(Prohibit modification and deletion)

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	res = make([]int, 0)
	preorder(root, &res)
	return res
}
func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
