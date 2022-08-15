package _08_sortedArrayToBST

//leetcode submit region begin(Prohibit modification and deletion)

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if left > right {
		return nil
	}

	mid := left + (right-left)>>1

	root := &TreeNode{Val: nums[mid]}
	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
