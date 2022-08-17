package _22_maxProfit

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

// [7, 1, 5, 3, 6, 4]
// ==> 7
//[day][      空仓],[      持股]
//[  0]          0,        -7,
//[  1]          0,        -1,
//[  2]          4,        -1,
//[  3]          4,         1,
//[  4]          7,         1,
//[  5]          7,         3,
func maxProfit(prices []int) int {
	// var dp [][][]int
	// dp[n][k][0/1]
	// dp[n][k][0] = max(dp[n-1][k][0],dp[n][k][1]+prices[n])
	// dp[n][k][1] = max(dp[n-1][k][1],dp[n][k][0]-prices[n])
	// n 表示交易天数
	// k 表示交易次数
	// 0/1 表示是否持有股票  0表示不持有，1表示持有

	sz := len(prices)
	if sz == 0 {
		return 0
	}
	dp := make([][2]int, sz)
	// base case

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < sz; i++ {
		// max(
		//		dp[i-1][0] ==>(上一天就不持有股票，最大利润的值，今天也不买入)
		//      dp[i-1][1] ==>(上一天就持有股票时，最大利润的值)
		//      dp[i-1][1] + prices[i] ==>(上一天就持有股票时，以当天的price卖出时，最大利润的值)
		//)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	printArr(dp)
	return dp[sz-1][0]
}
func printArr(dp [][2]int) {
	fmt.Printf("[%3v][%8v],[%8v]\n", "day", "空仓", "持股")
	for i := 0; i < len(dp); i++ {
		fmt.Printf("[%3v] ", i)
		arr := dp[i]
		for j := 0; j < len(arr); j++ {
			fmt.Printf("%10v,", arr[j])
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitMath(prices []int) int {
	res := 0

	if len(prices) == 0 {
		return res
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
