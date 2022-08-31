package _03_canCross

// canCross 403. 青蛙过河
// 动态规划
func canCross(stones []int) bool {
	n := len(stones)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// dp[i][k] 表示当前在第 i 个石头上，上一步跳了 k步的时候，是否能到达
	dp[0][0] = true
	// check传入的数据是否合法
	for i := 1; i < n; i++ {
		// 两个石头之间的距离比较大，不可能跳过去
		if stones[i]-stones[i-1] > i {
			return false
		}
	}

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			// 若上一步的位置为j，那么跳跃的距离为stones[i]-stones[j]
			// k表示上一步的跳跃的距离
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]

			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false

}
