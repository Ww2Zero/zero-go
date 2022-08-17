package _23_mergeKLists

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	//return mergeKListsByDivide(lists)
	return mergeKListsByMinHeap(lists)
}
func mergeKListsByDivide(lists []*ListNode) *ListNode {
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

func mergeKListsByMinHeap(lists []*ListNode) *ListNode {
	k := len(lists)
	h := new(minHeap)
	for i := 0; i < k; i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}

		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
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
