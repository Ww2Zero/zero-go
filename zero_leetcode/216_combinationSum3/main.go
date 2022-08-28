package _16_combinationSum3

// combinationSum3
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var track []int
	var backtrack func(k, n, start, sum int)

	backtrack = func(k, n, start, sum int) {
		// 剪枝
		if sum > n {
			return
		}
		if len(track) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, track)
				res = append(res, tmp)
			} else {
				return
			}
		}
		// 9-(k-len(track))+1 剩余的元素是否够k个
		for i := start; i <= 9-(k-len(track))+1; i++ {
			sum += i
			track = append(track, i)
			backtrack(k, n, i+1, sum)
			sum -= i
			track = track[:len(track)-1]
		}
	}
	backtrack(k, n, 1, 0)
	return res
}
