package _0_subsetsWithDup

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	var dfs func(nums []int, temp []int, start int)
	dfs = func(nums []int, temp []int, start int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			temp = append(temp, nums[i])
			dfs(nums, temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(nums, []int{}, 0)

	return res
}
