package _05_longestPalindrome

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// expandAroundCenter 以i为中心向两边扩散
func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func longestPalindrome2(s string) string {
	// dp[i][j] 表示s[i:j]是否是回文串
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	leftPoint, rightPoint := 0, 1
	// 初始化
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	// 从长度为2开始，判断是否是回文串
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			// 由 L 和 i 可以确定右边界，即 j - i + 1 = L 得
			j := l + i - 1
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == true && j-i+1 > rightPoint-leftPoint+1 {
				leftPoint = i
				rightPoint = j
			}
		}
	}
	return s[leftPoint : rightPoint+1]
}
