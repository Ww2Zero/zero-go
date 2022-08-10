package main

//leetcode submit region begin(Prohibit modification and deletion)

//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//
//示例 1：
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//示例 2：
//
//输入：n = 0
//输出：0
//示例 3：
//
//输入：n = 1
//输出：0

func countPrimesByBrute(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrim(i) {
			count++
		}
	}
	return count
}
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	for i := 2; i*i < n; i++ {
		if isPrim(i) {
			for j := i * i; j < n; j += i {
				arr[j] = 0
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		count += arr[i]
	}
	return count
}

func isPrim(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
