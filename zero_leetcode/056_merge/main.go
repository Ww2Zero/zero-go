package _56_merge

import "sort"

func merge(intervals [][]int) [][]int {
	// 对数组进行排序，按照每个数组的第一个元素进行排序，小的元素放在前面，大的元素放在后面
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	// prev 用于记录重合数组的边界 [0]表示左边界 [1]表示右边界
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		// 如果当前数组的左边界大于上一个数组的右边界，说明两个数组不重合
		if prev[1] < cur[0] {
			// 将上一个数组加入到结果中
			res = append(res, prev)
			// 将当前数组赋值给prev，用于下一次比较
			prev = cur
		} else {
			// 如果当前数组的左边界小于等于上一个数组的右边界，说明两个数组重合
			// 将上一个数组的右边界更新为上一个数组的右边界和当前数组的右边界的最大值
			prev[1] = max(prev[1], cur[1])
		}
	}
	// 将最后一个区间加入到结果中
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
