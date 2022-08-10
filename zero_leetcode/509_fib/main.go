package main

//斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
//F(0) = 0，F(1) = 1
//F(n) = F(n - 1) + F(n - 2)，其中 n > 1
//给定 n ，请计算 F(n) 。
//
//示例 1：
//
//输入：n = 2
//输出：1
//解释：F(2) = F(1) + F(0) = 1 + 0 = 1
//示例 2：
//
//输入：n = 3
//输出：2
//解释：F(3) = F(2) + F(1) = 1 + 1 = 2
//示例 3：
//
//输入：n = 4
//输出：3
//解释：F(4) = F(3) + F(2) = 2 + 1 = 3
//leetcode submit region begin(Prohibit modification and deletion)

func fib(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return n
	}
	res, l, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		res = l + r
		l = r
		r = res
	}
	return res
}

func fibBase(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return n
	}
	return fibBase(n-1) + fibBase(n-2)
}

func fibDP(n int) int {

	if n < 0 {
		return -1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 0 1 1 2 3 5
func fibMemo(n int) int {
	if n < 0 {
		return -1
	}
	// memo
	memo := make([]int, n+1)
	return memoHelper(memo, n)
}

func memoHelper(memo []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = memoHelper(memo, n-1) + memoHelper(memo, n-2)
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)
