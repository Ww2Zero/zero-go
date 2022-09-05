package _44_preorderTraversal

import "container/list"

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

func preorderTraversal2(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)
	var sts []*TreeNode
	sts = append(sts, root)
	for len(sts) > 0 {
		node := sts[len(sts)-1]
		ans = append(ans, node.Val)
		sts = sts[:len(sts)-1]
		if node.Right != nil {
			sts = append(sts, root.Right)
		}
		if node.Left != nil {
			sts = append(sts, root.Left)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
