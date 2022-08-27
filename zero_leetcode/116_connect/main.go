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

// connectByLevel 通过层序遍历的方式
func connectByLevel(root *Node) *Node {
	if root == nil {
		return nil
	}
	var res [][]*Node
	var q []*Node

	q = append(q, root)

	for len(q) > 0 {
		var p, c []*Node
		for i := 0; i < len(q); i++ {
			node := q[i]
			c = append(c, node)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		res = append(res, c)
		q = p
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i])-1; j++ {
			res[i][j].Next = res[i][j+1]
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
