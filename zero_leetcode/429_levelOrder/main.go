package _29_levelOrder

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 429 N叉树的层序遍历
func levelOrder(root *Node) [][]int {
	var res [][]int
	var q []*Node
	if root == nil {
		return res
	}
	q = append(q, root)
	for len(q) > 0 {
		var subRes []int
		var p []*Node
		for i := 0; i < len(q); i++ {
			sub := q[i]
			subRes = append(subRes, sub.Val)
			for j := 0; j < len(sub.Children); j++ {
				p = append(p, sub.Children[j])
			}
		}
		res = append(res, subRes)
		q = p
	}
	return res
}
