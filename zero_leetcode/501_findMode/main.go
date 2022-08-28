package _01_findMode

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	cnt := 1
	max := 1

	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			cnt++
		} else {
			cnt = 1
		}

		if cnt >= max {
			if cnt > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = cnt
		}

		prev = node
		travel(node.Right)
	}

	travel(root)

	return res
}
