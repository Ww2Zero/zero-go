package _55_canJump

import "fmt"

// 55. 跳跃游戏
// dp[i]表示从0到i的最大跳路数
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				fmt.Println(dp)
				break
			}
		}
	}
	return dp[len(nums)-1]
}

// canJumpByGreedy 贪心算法
func canJumpByGreedy(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	cover := 0
	for i := 0; i <= cover; i++ {
		// cover是最远能到达的位置范围
		// nums[i]是当前位置的跳跃距离
		// i 表示当前所在的位置
		// nums[i] + i 表示当前可以能到达的最远位置
		cover = max(cover, nums[i]+i)
		// 如果cover已经超过了最后一个位置，说明可以到达最后一个位置
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
