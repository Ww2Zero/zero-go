package _23_maxProfit

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	sz := len(prices)
	if sz == 0 {
		return 0
	}
	dp := make([][3][2]int, sz)
	// base case
	for i := 0; i < sz; i++ {
		for k := 2; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[sz-1][2][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
