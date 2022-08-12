package _16_connect

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoHelper(root.Left, root.Right)
	return root
}

func connectTwoHelper(n1, n2 *Node) {
	if n1 == nil {
		return
	}

	n1.Next = n2

	connectTwoHelper(n1.Left, n1.Right)
	connectTwoHelper(n2.Left, n2.Right)
	connectTwoHelper(n1.Right, n2.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
