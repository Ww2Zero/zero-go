package _05_buildTree

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree 根据tree前序遍历和中序遍历构建tree
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 用map存储value -> index  的数据，避免在inorder中查找rootNode时候需要遍历
	valToIndexMap := make(map[int]int)
	for i, v := range inorder {
		valToIndexMap[v] = i
	}
	return buildHelper2(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1,
		valToIndexMap,
	)
}

// buildTreeBase 基础的版本，未优化
func buildTreeBase(preorder []int, inorder []int) *TreeNode {
	return buildHelper(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

func buildHelper2(preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int,
	valToIndexMap map[int]int,
) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	rootIdx := valToIndexMap[rootVal]

	lenSize := rootIdx - inStart

	root.Left = buildHelper2(preorder, preStart+1, preStart+lenSize,
		inorder, inStart, rootIdx-1, valToIndexMap)
	root.Right = buildHelper2(preorder, preStart+lenSize+1, preEnd,
		inorder, rootIdx+1, inEnd, valToIndexMap)
	return root
}
func buildHelper(
	preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int,
) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	// find inorder left tree nodes
	rootIdx := -1
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIdx = i
		}
	}
	lenLeft := rootIdx - inStart

	root.Left = buildHelper(preorder, preStart+1, preStart+lenLeft,
		inorder, inStart, rootIdx-1)
	root.Right = buildHelper(preorder, preStart+lenLeft+1, preEnd,
		inorder, rootIdx+1, inEnd)

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
