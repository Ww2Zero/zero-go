package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	hasCycle, meetNode := isCycle(head)
	if !hasCycle {
		return nil
	}
	fast := head
	slow := meetNode
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
func isCycle(head *ListNode) (hasCycle bool, meetNode *ListNode) {
	fast := head
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true, fast
		}
	}
	return false, nil
}
