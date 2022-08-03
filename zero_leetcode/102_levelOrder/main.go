package main

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		l := []int{}
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			l = append(l, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		ret = append(ret, l)
		q = p
	}
	return ret

}
