package _23_mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		cur := lists[i]
		res = mergeTwoLists(res, cur)
	}
	return res
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
