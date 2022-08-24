package _09_minSubArrayLen

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	res := math.MaxInt
	sum := 0
	j := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		// TODO WatchMe: for sum >= target WHY?
		for sum >= target {
			subLen := i - j + 1
			if subLen < res {
				res = subLen
			}
			sum -= nums[j]
			j++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
