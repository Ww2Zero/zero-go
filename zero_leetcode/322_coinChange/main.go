package _22_coinChange

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

// 输入 [1,2,5] 11
// 输出 3
func coinChangeBrute(coins []int, amount int) (res int) {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res = math.MaxInt
	for i := 0; i < len(coins); i++ {
		cur := coins[i]
		subProblem := coinChangeBrute(coins, amount-cur)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	if res == math.MinInt {
		return -1
	}
	return res
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
func coinChange(coins []int, amount int) (res int) {
	memo := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		memo[i] = -666
	}
	return coinChangeMemo(coins, amount, memo)
}
func coinChangeMemo(coins []int, amount int, memo []int) (res int) {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != -666 {
		return memo[amount]
	}
	res = math.MaxInt
	for i := 0; i < len(coins); i++ {
		cur := coins[i]
		subProblem := coinChangeMemo(coins, amount-cur, memo)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	if res == math.MaxInt {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

//leetcode submit region end(Prohibit modification and deletion)
