package _60_getPermutation

import "fmt"

//给出集合[1,2,3,...,n]，其所有元素共有n! 种排列。
//
//按大小顺序列出所有排列情况，并一一标记，当n = 3 时, 所有排列如下：
//
//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//给定n 和k，返回第k个排列。
// getPermutation 60. 第k个排列
// 回溯算法
func getPermutation(n int, k int) string {
	var track []int
	var res string
	cnt := 0
	isUsed := make([]bool, n+1)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level > n {
			return
		}
		if len(track) == n {
			fmt.Println(track)
			cnt++
			if cnt == k {
				for _, v := range track {
					res += string(rune(v + '0'))
				}
			}
			return
		}

		for i := 1; i <= n; i++ {
			if isUsed[i] { // 跳过已使用的数
				continue
			}
			isUsed[i] = true
			track = append(track, i)
			backtracking(level + 1)
			track = track[:len(track)-1]
			isUsed[i] = false
		}
	}

	backtracking(0)
	return res
}
