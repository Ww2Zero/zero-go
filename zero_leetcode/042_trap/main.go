package _42_trap

//leetcode submit region begin(Prohibit modification and deletion)

// 	// water[i] = min(max(height[0...i]),max(height[i...end]))-height[i]
func trap(height []int) int {
	sz := len(height)
	if sz == 0 {
		return 0
	}
	res := 0
	// lMax, 记录从左到右每个位置中经过point的最高点
	// rMax, 记录从右到左每个位置中经过point的最高点
	lMax, rMax := make([]int, sz), make([]int, sz)

	// init left->right leftMax Array
	lMax[0] = height[0]
	for i := 1; i < sz; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}
	// init right->left rightMax Array
	rMax[sz-1] = height[sz-1]
	for i := sz - 2; i >= 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		res += min(lMax[i], rMax[i]) - height[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
