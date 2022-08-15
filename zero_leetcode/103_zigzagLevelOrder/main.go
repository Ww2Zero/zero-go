package _03_zigzagLevelOrder

//leetcode submit region begin(Prohibit modification and deletion)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	flag := true
	for len(q) > 0 {
		var l []int
		var p []*TreeNode
		// 通过flag来控制本层中数据遍历顺序，从前到后，还是从后到前
		if flag {
			for j := 0; j < len(q); j++ {
				node := q[j]
				l = append(l, node.Val)
			}
		} else {
			for j := len(q) - 1; j >= 0; j-- {
				node := q[j]
				l = append(l, node.Val)
			}
		}
		// 节点的添加顺序不变，依旧是从左到右
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		ret = append(ret, l)
		q = p
		flag = !flag
	}
	return ret

}

//leetcode submit region end(Prohibit modification and deletion)
