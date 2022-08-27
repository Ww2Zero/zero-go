package _26_invertTree

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// invertTreeByLevel
//Time complexity: O(n)
func invertTreeByLevel(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var p []*TreeNode
		for i := 0; i < len(q); i++ {
			node := q[i]
			node.Left, node.Right = node.Right, node.Left

			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
