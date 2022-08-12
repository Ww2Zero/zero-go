package _50_deleteNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// del
		// case 1 当前节点是叶子节点 无子节点
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// case 2 当前节点有一个叶子节点
		// 只有一个右节点
		if root.Left == nil {
			return root.Right
		}
		// 只有一个左节点
		if root.Right == nil {
			return root.Left
		}
		// case 3 当前节点有2个叶子节点
		// 获取当前的右节点上的最小节点
		min := getMinNode(root.Right)
		// 替换两个节点的节点值
		root.Val = min.Val
		// 删除原本的节点的值
		root.Right = deleteNode(root.Right, min.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMinNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
