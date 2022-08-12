package _0_climbStairs

//leetcode submit region begin(Prohibit modification and deletion)

func climbStairs(n int) int {
	//return climbStairsBase(n)
	//return climbStairsMemo(n)
	//return climbStairsDP(n)
	return climbStairsBest(n)
}

// climbStairsBest 最优解法
func climbStairsBest(n int) int {
	if n < 0 {
		return -1
	}
	if n == 1 || n == 2 {
		return n
	}
	r, lo, hi := 0, 1, 2
	for i := 3; i <= n; i++ {
		r = lo + hi
		lo = hi
		hi = r
	}
	return r
}

//climbStairsDP  dpTable 迭代
func climbStairsDP(n int) int {
	if n < 0 {
		return -1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//climbStairsMemo 备忘录模式 递归
func climbStairsMemo(n int) int {
	if n < 0 {
		return -1
	}
	memo := make([]int, n+1)
	return climbStairsHelper(memo, n)
}

// 分析的基础过程
// fn 有两种上楼的方式 f(n-1) 和 f(n-2),那么上楼的方式为f(n-1)+ f(n-2)
// n=1  1                                1
// n=2  1+1; 2                           2
// n=3  1+1+1;2+1;1+2                    3
// n=4  1+1+1+1;2+2;2+1+1;1+1+2;1+2+1    5
func climbStairsBase(n int) int {
	if n < 0 {
		return -1
	}
	if n == 1 || n == 2 {
		return n
	}
	return climbStairsBase(n-1) + climbStairsBase(n-2)
}

func climbStairsHelper(memo []int, n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = climbStairsHelper(memo, n-1) + climbStairsHelper(memo, n-2)
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)
