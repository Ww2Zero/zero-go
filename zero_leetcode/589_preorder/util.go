package main

import "math"

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

//NewTree [1,null,3,2,4,null,5,6]
func NewTree(ins []int) (root *Node) {
	// todo fix
	if len(ins) == 0 {
		return nil
	}
	root = &Node{Val: ins[0]}
	for i := 1; i < len(ins); i++ {
		if ins[i] == math.MinInt {
			continue
		}
		root.Children = append(root.Children, NewTree(ins[i:]))
		break
	}

	return
}
