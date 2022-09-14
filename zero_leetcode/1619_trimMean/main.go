package _619_trimMean

import "sort"

// trimMean 1619. 删除某些元素后的数组均值
func trimMean(arr []int) float64 {
	sort.Ints(arr)

	sz := len(arr)
	sum := 0

	for _, v := range arr[sz/20 : sz*19/20] {
		sum += v
	}

	return float64(sum*10) / float64(sz*9)
}
