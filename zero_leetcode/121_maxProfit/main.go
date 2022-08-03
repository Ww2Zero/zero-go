package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	cost := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		cost = min(cost, prices[i])
		profit = max(profit, prices[i]-cost)
	}
	return profit

}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
