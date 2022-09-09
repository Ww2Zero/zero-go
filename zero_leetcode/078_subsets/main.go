package _78_subsets

import "sort"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	var backtracking func(nums []int, tmp []int, start int)

	backtracking = func(nums []int, tmp []int, start int) {
		if start > len(nums) {
			return
		}
		temp := make([]int, len(tmp))
		copy(temp, tmp)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			// 必须从i+1开始，因为不能重复使用
			backtracking(nums, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backtracking(nums, []int{}, 0)
	return res
}
