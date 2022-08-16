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
	return max(robHelper(nums, 0, l-2, memo1), robHelper(nums, 1, l-1, memo2))
}

func robHelper(nums []int, start, end int, memo []int) int {
	if start > end {
		return 0
	}

	if memo[start] != -1 {
		return memo[start]
	}
	memo[start] = max(
		robHelper(nums, start+2, end, memo)+nums[start],
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
