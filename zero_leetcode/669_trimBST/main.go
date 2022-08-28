package _69_trimBST

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// trimBST 669. 修剪二叉搜索树
// 利用二叉搜索树的特性，修剪二叉搜索树
// 修剪后的二叉搜索树仍然是二叉搜索树
// 修剪后的二叉搜索树的节点值都在[L, R]之间
// 递归
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 当前节点值小于low，修剪左子树
	// 如果root（当前节点）的元素小于low的数值，那么应该递归右子树，并返回右子树符合条件的头结点。
	if root.Val < low {
		right := trimBST(root.Right, low, high)
		return right
	}
	// 当前节点值大于high，修剪右子树
	// 如果root（当前节点）的元素大于high的数值，那么应该递归左子树，并返回左子树符合条件的头结点。
	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 迭代

func trimBST2(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 处理 root，让 root 移动到[low, high] 范围内，注意是左闭右闭
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	// 此时 root 已经在[low, high] 范围内，处理左孩子元素小于 low 的情况（左节点是一定小于 root.Val，因此天然小于 high）
	cur1 := root
	for cur1 != nil {
		for cur1.Left != nil && cur1.Left.Val < low {
			cur1.Left = cur1.Left.Right
		}
		cur1 = cur1.Left
	}
	// 此时 root 已经在[low, high] 范围内，处理右孩子大于 high 的情况
	cur2 := root
	for cur2 != nil {
		for cur2.Right != nil && cur2.Right.Val > high {
			cur2.Right = cur2.Right.Left
		}
		cur2 = cur2.Right
	}

	return root
}
