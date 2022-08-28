package _40_combinationSum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	// 先排序.排序后,
	sort.Ints(candidates)
	var backtrack func(sum, start int)
	backtrack = func(sum, start int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			// 剪枝，判断若是相同的数字，则跳过
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			track = append(track, candidates[i])
			backtrack(sum, i+1)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0, 0)
	return res
}
