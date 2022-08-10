package main

//leetcode submit region begin(Prohibit modification and deletion)
//给定两个用链表表示的整数，每个节点包含一个数位。
//
//这些数位是反向存放的，也就是个位排在链表首部。
//
//编写函数对这两个整数求和，并用链表形式返回结果。
//
//示例：
//
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
// 输出：2 -> 1 -> 9，即912
// (7 -> 1 -> 6)
// (5 -> 9 -> 2)
// ---------------
// 2 ->1+9+1->1+6+2

//输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	pre := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		pre.Next = &ListNode{Val: carry % 10}
		pre = pre.Next
		carry = carry / 10
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
