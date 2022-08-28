package _77_combine

// combine 77. 组合
// 回溯算法
func combine(n int, k int) [][]int {
	var res [][]int
	if n < 0 || k < 0 || k > n {
		return res
	}
	var track []int
	var backtrack func(n, k, start int)
	backtrack = func(n, k, start int) {
		if len(track) == k {
			// 注意这里要拷贝一份
			t := make([]int, k)
			copy(t, track)
			res = append(res, t)
		}
		// 剪枝 判断剩余的元素是否够k个
		if len(track)+(n-start+1) < k {
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(n, k, i+1)
			track = track[:len(track)-1]
		}
	}
	backtrack(n, k, 1)
	return res
}
