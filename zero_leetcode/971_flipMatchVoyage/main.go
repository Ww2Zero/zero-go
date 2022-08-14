package _71_flipMatchVoyage

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [1,2,3]
//  [1,3,2]
//   --------> [1]
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	return flipMatchVoyage2(root, voyage) // right
	//return flipMatchVoyageexecute(root, voyage) // right
	//return flipMatchVoyageDFS(root, voyage) // right
	//return flipMatchVoyageDFS2(root, voyage)
}

func flipMatchVoyageDFS2(root *TreeNode, voyage []int) []int {
	n := 0
	ans := []int{}
	var dfs func(*TreeNode, []int) bool
	dfs = func(node *TreeNode, v []int) bool {
		if root == nil {
			return true
		}
		if root.Val != v[n] {
			return false
		}
		n = n + 1
		if root.Left != nil && root.Left.Val != v[n+1] {
			ans = append(ans, root.Val)
			return dfs(root.Right, v) && dfs(root.Left, v)
		}
		return dfs(root.Left, v) && dfs(root.Right, v)
	}
	if dfs(root, voyage) {
		return ans
	}
	return []int{-1}
}

func flipMatchVoyageDFS(root *TreeNode, voyage []int) []int {
	n := len(voyage)
	ans := []int{}
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, cur int) int {
		if node != nil && node.Val == voyage[cur] {
			cur++
			if node.Left != nil && node.Left.Val == voyage[cur] {
				ldfs := dfs(node.Left, cur)
				return dfs(node.Right, ldfs)
			} else if node.Right != nil && node.Right.Val == voyage[cur] {
				if node.Left != nil {
					ans = append(ans, node.Val)
				}
				rdfs := dfs(node.Right, cur)
				return dfs(node.Left, rdfs)
			}
		}
		return cur
	}
	if n == dfs(root, 0) {
		return ans
	}
	return []int{-1}
}

func flipMatchVoyagedfs(root *TreeNode, voyage []int) []int {
	var res []int
	dfs(root, voyage, &res, 0)
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyageexecute(root *TreeNode, voyage []int) []int {
	if nil == root {
		return []int{-1}
	}
	var ans []int
	index := 0
	flag := execute(root, voyage, &index, &ans)
	if !flag || index < len(voyage)-1 { //如果不符合要求或者数组长度大于树节点总数
		return []int{-1}
	}
	return ans
}

func execute(root *TreeNode, voyage []int, index *int, ans *[]int) bool {
	if nil == root { //如果是空节点，则将索引回退
		*index -= 1
		return true
	}
	if *index >= len(voyage) { //如果索引大于数组长度，并且当前节点还不为空，则树长度大于数组长度，不符合要求
		return false
	}
	if root.Val != voyage[*index] { //如果当前节点不等于索引指向数组的值，则不符合要求
		*index -= 1
		return false
	}

	*index += 1
	left := execute(root.Left, voyage, index, ans)
	if !left { //左子树不符合要求
		flip(root) //翻转
		*index += 1
		left := execute(root.Left, voyage, index, ans) //再判断左子树
		if !left {                                     //还不符合，则返回false
			return false
		}
		*ans = append(*ans, root.Val) //符合要求，将节点值加入结果中
	}
	*index += 1
	right := execute(root.Right, voyage, index, ans) //判断右子树
	if !right {
		return false
	}
	return true
}

func flip(root *TreeNode) {
	if nil == root {
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func flipMatchVoyage2(root *TreeNode, voyage []int) []int {
	var res []int
	dfs2(root, &voyage, &res)
	return res
}

// 使用res记录需要反转的节点的值
func dfs2(root *TreeNode, voyage *[]int, res *[]int) bool {
	if root == nil {
		return true
	}
	// 根节点的值和voyage第0个元素不符，无法通过反转节点使使先序遍历符合voyage
	if root.Val != (*voyage)[0] {
		*res = []int{-1}
		return false
	}
	// 当且仅当根节点的左右节点都不为空，且右儿子的值等于voyage的第1个元素，需要反转根节点的左右儿子
	if root.Left != nil && root.Right != nil && root.Right.Val == (*voyage)[1] {
		*res = append(*res, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}
	// 消耗掉voyage第0个元素
	*voyage = (*voyage)[1:]
	// 如果对左儿子递归的结果为false，就进行剪枝操作
	ok := dfs2(root.Left, voyage, res)
	if !ok {
		return false
	}
	return dfs2(root.Right, voyage, res)
}

func dfs(root *TreeNode, voyage []int, res *[]int, idx int) bool {
	if root == nil {
		return true
	}
	if root.Val != voyage[idx] {
		*res = []int{-1}
		return false
	}
	if root.Left != nil && root.Left.Val != voyage[idx+1] {
		*res = append(*res, root.Val)
		root.Right, root.Left = root.Left, root.Right
	}
	idx++
	ok := dfs(root.Left, voyage, res, idx)
	if !ok {
		return false
	}
	idx++
	return dfs(root.Right, voyage, res, idx)
}
