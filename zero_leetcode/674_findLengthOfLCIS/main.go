package _74_findLengthOfLCIS

//leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {

	sz := len(nums)
	if sz == 0 {
		return 0
	}
	dp := make([]int, sz)

	for i := 0; i < sz; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < sz; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
