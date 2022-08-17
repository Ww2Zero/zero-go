package _23_maxProfit

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	sz := len(prices)
	if sz == 0 {
		return 0
	}
	dp := make([][3][2]int, sz)
	// base case
	for i := 0; i < sz; i++ {
		for k := 2; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	printArr(dp)
	return dp[sz-1][2][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printArr(dp [][3][2]int) {

}
func printArr2(dp [][2]int) {
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

//leetcode submit region end(Prohibit modification and deletion)
