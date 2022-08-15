package _09_sortedListToBST

//leetcode submit region begin(Prohibit modification and deletion)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return build(head, nil)
}

func build(start, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}
	mid := getMid(start, end)
	root := &TreeNode{Val: mid.Val}
	root.Left = build(start, mid)
	root.Right = build(mid.Next, end)
	return root
}

func getMid(start, end *ListNode) *ListNode {
	fast, slow := start, start

	for fast != nil && slow != nil && fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
