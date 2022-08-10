package main

import "fmt"

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
func helpLog(head *ListNode) func() {
	helpF(head, "start")
	return func() {
		helpF(head, "exit")
	}
}

func helpF(head *ListNode, str string) {
	fmt.Printf("[%12s]=> ", str)
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Println("Nil")
}
