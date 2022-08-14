package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		// step1 找到最小的价格
		minPrice = min(minPrice, prices[i])
		// step2 然后获取最大的价格
		profit = max(profit, prices[i]-minPrice)
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
