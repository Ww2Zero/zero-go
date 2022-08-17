package _15_threeSum

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

func nSumTarget(nums []int, n, start, target int) [][]int {
	var res [][]int
	sz := len(nums)
	if n < 2 || sz < n {
		return res
	}
	if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			left, right := nums[lo], nums[hi]
			sum := left + right
			if sum > target {
				for lo < hi && right == nums[hi] {
					hi--
				}
			} else if sum < target {
				for lo < hi && left == nums[lo] {
					lo++
				}
			} else {
				res = append(res, []int{left, right})
				for lo < hi && left == nums[lo] {
					lo++
				}
				for lo < hi && right == nums[hi] {
					hi--
				}
			}
		}
	} else {
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for j := 0; j < len(sub); j++ {
				sub[j] = append(sub[j], nums[i])
				res = append(res, sub[j])
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
