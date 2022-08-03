package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			dummy.Next = list2
			list2 = list2.Next
		} else {
			dummy.Next = list1
			list1 = list1.Next
		}
		dummy = dummy.Next
	}

	if list1 != nil {
		dummy.Next = list1
	}
	if list2 != nil {
		dummy.Next = list2
	}

	return head.Next
}
