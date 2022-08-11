package main

func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode
	for head != nil {

		next := head.Next

		head.Next = dummy

		dummy = head
		head = next
	}
	return dummy
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//func reverseList(head *ListNode) *ListNode {
//	log.Printf("reverseList before ,head=%v %v", head.Val, head.Next)
//	if head.Next == nil {
//		return head
//	}
//	r := reverseList(head.Next)
//	log.Printf("reverseList after,r=%v %v", r.Val, r.Next)
//	log.Printf("reverseList after,head=%v %v", head.Val, head.Next)
//	r.Next = head
//	log.Printf("reverseList after,r=%v %v", r.Val, r.Next)
//	return r.Next
//}
