package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
func reverseKGroup(head *ListNode, k int) *ListNode {

	left, right := head, head
	// base case
	// if len(head) < k .then return head node
	for i := 0; i < k; i++ {
		if right == nil {
			return left
		}
		right = right.Next
	}

	// move node a and node b
	// reverse [a,b)
	dummy := reverseBetween(left, right)
	// move head node to next need reverse node head point
	head = right
	// recursion solution
	left.Next = reverseKGroup(head, k)
	// return node
	return dummy
}

func reverseBetween2(left, right *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = left
	pre := dummy

	for pre != right {
		next := pre.Next

		pre.Next = next.Next

		next.Next = pre.Next

		pre = next.Next
	}
	return dummy.Next
}
func reverseBetween(left, right *ListNode) *ListNode {
	var dummy *ListNode
	helpF(left, "left")
	helpF(right, "right")

	for left != right {
		fmt.Println("===========")
		// 1. 记录left的下一个节点next
		next := left.Next
		helpF(next, "next")
		// 2. 修改原本的left.Next节点，
		//  当dummy为null时，则dummy为
		helpF(dummy, "dummy")
		left.Next = dummy
		helpF(left, "left")
		helpF(left.Next, "left.Next")
		dummy = left
		left = next
	}
	return dummy
}
