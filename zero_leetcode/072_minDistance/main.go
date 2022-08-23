package _72_minDistance

//leetcode submit region begin(Prohibit modification and deletion)
// 编辑距离
// https://leetcode.cn/problems/edit-distance/solution/dai-ma-sui-xiang-lu-72-bian-ji-ju-chi-do-y87e/
func minDistance(word1 string, word2 string) int {
	sz1, sz2 := len(word1), len(word2)
	dp := make([][]int, sz1+1)
	for i := 0; i < sz1+1; i++ {
		dp[i] = make([]int, sz2+1)
	}
	for i := 0; i < sz1+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < sz2+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < sz1+1; i++ {
		for j := 1; j < sz2+1; j++ {
			// 相同时，i指针和j指针可以直接移动，无需操作，所以dp[i][j]的修改次数等于dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//
				dp[i][j] = min(
					// 增加
					dp[i][j-1]+1,
					// 删除
					dp[i-1][j]+1,
					// 替换
					dp[i-1][j-1]+1)
			}
		}
	}
	return dp[sz1][sz2]
}

func min(a, b, c int) int {
	return min2(min2(a, b), c)
}
func min2(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
