package _04_sumOfLeftLeaves

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sumOfLeftLeaves 404. 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftValue := sumOfLeftLeaves(root.Left)
	rightValue := sumOfLeftLeaves(root.Right)

	midValue := 0

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midValue = root.Left.Val
	}
	return midValue + leftValue + rightValue
}
