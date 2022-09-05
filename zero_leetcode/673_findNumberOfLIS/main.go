package _73_findNumberOfLIS

import "fmt"

//给定一个未排序的整数数组，找到最长递增子序列的个数。
//
//示例 1:
//
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//示例 2:
//
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。

// findNumberOfLIS 动态规划
func findNumberOfLIS(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	//dp[i]：i之前（包括i）最长递增子序列的长度为dp[i]
	//count[i]：以nums[i]为结尾的字符串，最长递增子序列的个数为count[i]
	dp := make([]int, size)
	for i, _ := range dp {
		dp[i] = 1
	}
	count := make([]int, size)
	for i, _ := range count {
		count[i] = 1
	}

	maxCount := 0
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					// 长度增加，数量不变 count[i] = count[j]
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					// 最长递增子序列的长度并没有增加，但是出现了长度一样的情况，数量增加 count[i] += count[j]
					count[i] += count[j]
				}
				dp[i] = max(dp[j]+1, dp[i])
			}
			maxCount = max(maxCount, dp[i])
		}
		fmt.Println(dp)
		fmt.Println(count)
		fmt.Println("====")
	}

	result := 0
	for i := 0; i < size; i++ {
		if maxCount == dp[i] {
			result += count[i]
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
