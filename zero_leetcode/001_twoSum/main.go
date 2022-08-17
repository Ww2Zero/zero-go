package _01_twoSum

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if idx, ok := m[nums[i]]; ok {
			return []int{i, idx}
		} else {
			m[target-nums[i]] = i
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
