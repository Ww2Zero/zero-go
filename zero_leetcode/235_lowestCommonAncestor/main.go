package _35_lowestCommonAncestor

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor 35. 求二叉搜索树的最近公共祖先
// 利用二叉搜索树的特性，求出两个节点的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res := root
	for {
		if res.Val > p.Val && res.Val > q.Val {
			res = res.Left
		} else if res.Val < p.Val && res.Val < q.Val {
			res = res.Right
		} else {
			return res
		}

	}
	return res
}
