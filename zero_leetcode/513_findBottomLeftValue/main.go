package _13_findBottomLeftValue

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var maxdeep int
	var maxdeepLeftVal int
	var findLeftValue func(root *TreeNode, deep int)
	findLeftValue = func(root *TreeNode, deep int) {
		if root.Left == nil && root.Right == nil {
			if deep > maxdeep {
				maxdeepLeftVal = root.Val
				maxdeep = deep
			}
		}
		if root.Left != nil {
			deep++
			findLeftValue(root.Left, deep)
			deep--
		}
		if root.Right != nil {
			deep++
			findLeftValue(root.Right, deep)
			deep--
		}
	}

	findLeftValue(root, 0)
	return maxdeepLeftVal
}
