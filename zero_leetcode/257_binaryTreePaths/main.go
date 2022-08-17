package _57_binaryTreePaths

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	binaryTreePathsHelper(root, "")
	return paths
}

func binaryTreePathsHelper(root *TreeNode, path string) {

	if root == nil {
		return
	}
	pathBuf := path
	pathBuf += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		paths = append(paths, pathBuf)
	} else {
		pathBuf += "->"
		binaryTreePathsHelper(root.Left, pathBuf)
		binaryTreePathsHelper(root.Right, pathBuf)
	}

}

//leetcode submit region end(Prohibit modification and deletion)
