package _76_wiggleMaxLength

// wiggleMaxLength 376. 摆动序列
// 贪心算法
func wiggleMaxLength(nums []int) int {
	var preDiff, curDiff, res int

	res = 1

	if len(nums) < 2 {
		return res
	}
	for i := 1; i < len(nums); i++ {
		curDiff = nums[i] - nums[i-1]
		//
		if curDiff > 0 && preDiff <= 0 {
			preDiff = curDiff
			res++
		}
		if curDiff < 0 && preDiff >= 0 {
			preDiff = curDiff
			res++
		}
	}
	return res
}

// wiggleMaxLengthDP 376. 摆动序列
// DP
func wiggleMaxLengthDP(nums []int) int {
	dp := make([][]int, len(nums))
	for i, _ := range dp {
		// dp[i][0] 表示i为波峰的时候，摆动序列的最大长度
		// dp[i][1] 表示i为波谷的时候，摆动序列的最大长度
		dp[i] = make([]int, 2)
	}
	// base case
	for i, _ := range dp {
		// i 自己可以成为波峰或者波谷，所以dp[i][0] = 1
		dp[i][0] = 1
		dp[i][1] = 1
	}

	for i := 1; i < len(nums); i++ {
		// 从0开始计算到i-1的最大的波峰和波谷的长度
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				//dp[i][0] = max(当前i的摇摆长度，当前j为波谷的时候摆动序列的最大长度+1)
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			} else {
				//dp[i][1] = max(当前i的摇摆长度，当前j为波峰的时候摆动序列的最大长度+1)
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
