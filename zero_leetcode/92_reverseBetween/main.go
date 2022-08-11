package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//示例 2：
//输入：head = [5], left = 1, right = 1
//输出：[5]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	pre := dummy

	//1. move pre point to left-1
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	head = pre.Next
	helpF(head, "head->")
	helpF(pre, "pre->")
	//2. reverse right-left list
	//for i := left; i < right; i++ {
	//	next := head.Next
	//
	//	head.Next = next.Next
	//
	//	next.Next = pre.Next
	//
	//	pre.Next = next
	//
	//}
	for i := left; i < right; i++ {
		next := head.Next

		head.Next = next.Next

		next.Next = pre.Next

		pre.Next = next
	}
	//3. return
	return dummy.Next
}
