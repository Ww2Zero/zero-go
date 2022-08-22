package _3_maxSubArray

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region end(Prohibit modification and deletion)
