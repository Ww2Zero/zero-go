package main

func preorder(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node != nil {
			ans = append(ans, node.Val)
			for _, child := range node.Children {
				dfs(child)
			}
		}
	}
	dfs(root)
	return
}
