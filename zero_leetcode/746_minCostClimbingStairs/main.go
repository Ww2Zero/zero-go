package _46_minCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	// dpTable中记录的数据 索引i 记录值v
	// 表示到达位置i 需要花费的最小金额v
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		// c1 表示爬一步上来的i-1 的最小花费 dp[i-1]+cost[i-1]
		c1 := dp[i-1] + cost[i-1]
		// c2 表示爬两步上来的i-2 的最小花费 dp[i-2]+cost[i-2]
		c2 := dp[i-2] + cost[i-2]
		// dp[i] 获取i-1或 i-2 的最小花费
		dp[i] = min(c1, c2)
	}
	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
