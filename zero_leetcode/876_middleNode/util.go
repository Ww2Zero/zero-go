package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(ints ...int) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for _, i := range ints {
		dummy.Next = &ListNode{Val: i}
		dummy = dummy.Next
	}
	return head.Next
}
