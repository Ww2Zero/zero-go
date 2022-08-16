package _13_rob

//leetcode submit region begin(Prohibit modification and deletion)

// rob 打家劫舍不一定是间隔一家打击一家，而是可能出现，选择价值高的家庭，不管中间间隔多少家。例如【1，2，55，2，3，66，2】，需要选择55，66，才可能是最大的价值
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	memo1, memo2 := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		memo1[i] = -1
		memo2[i] = -1
	}
	// 首位相连，
	// 1，选择从0开始，之后不能选择l-1。只能是l-2了
	// 2。选择从1开始，则可以选择到最后一个节点。l-1
	return max(robHelper(nums, 0, l-2, memo1), robHelper(nums, 1, l-1, memo2))
}

func robHelper(nums []int, start, end int, memo []int) int {
	if start > end {
		return 0
	}

	if memo[start] != -1 {
		return memo[start]
	}
	// memo表示memo[start]的最大值
	// 当前选择了当前节点之后，不能选择start+1的节点，所以最大的价值为当前节点的价值+start+2的最大价值
	memo[start] = max(
		// 选择当前节点时的最大价值，即当前节点【start】 + 【start+2】
		nums[start]+robHelper(nums, start+2, end, memo),
		// 不选择当前节点，即选择节点【start+1】
		robHelper(nums, start+1, end, memo),
	)
	return memo[start]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
