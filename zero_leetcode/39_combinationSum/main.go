package _9_combinationSum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int

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
			sum += candidates[i]
			track = append(track, candidates[i])
			backtrack(sum, i)
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0, 0)
	return res
}
