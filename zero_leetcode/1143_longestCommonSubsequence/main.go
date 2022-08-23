package _143_longestCommonSubsequence

func longestCommonSubsequence(text1 string, text2 string) int {

	n, m := len(text1), len(text2)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
