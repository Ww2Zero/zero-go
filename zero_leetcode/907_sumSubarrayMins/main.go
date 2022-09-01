package _07_sumSubarrayMins

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
// dp+单调栈
func sumSubarrayMins(A []int) int {
	stack, dp, res, mod := []int{}, make([]int, len(A)+1), 0, 1000000007
	stack = append(stack, -1)

	for i := 0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*A[i]) % mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
		fmt.Println(stack)
		//fmt.Println(dp)
	}
	return res
}

func sumSubarrayMins2(arr []int) int {
	stack := []int{}
	// 加入0 保证整个数组中的元素都能够弹出，最后的高度为0
	arr = append(arr, 0)
	res := 0
	if len(arr) == 0 {
		return res
	}
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevTopIndex := -1
			if len(stack) > 0 {
				prevTopIndex = stack[len(stack)-1]
			}
			left := topIndex - prevTopIndex - 1
			right := i - topIndex - 1
			res += ((1 + left + right + (right * left)) * arr[topIndex]) % 1000000007
			res = res % 1000000007
		}
		stack = append(stack, i)
	}
	return res
}

// 单调栈
func sumSubarrayMinsByStack(arr []int) int {
	// 单调递增栈
	stack := []int{}
	// 加入0 保证整个数组中的元素都能够弹出
	arr = append(arr, 0)
	// 保存结果的res
	res := 0
	// 遍历数组
	for i := 0; i < len(arr); i++ {
		// 如果栈不为空 并且当前元素小于栈顶元素
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			// 得到栈顶元素
			index := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			// 设定prevIndex为-1 栈为空时 设定 left 左边元素个数为0
			prevIndex := -1
			if len(stack) > 0 {
				// 获取下一个元素的位置
				prevIndex = stack[len(stack)-1]
			}
			// 得到以arr[index]为最小值的左边元素的个数和右边元素的个数
			// 通过（left + 1）(right + 1) 得到最终有多少个数组以该值作为最小值的个数
			// +1是包含该元素 不然会出现0的情况
			fmt.Println(prevIndex, index)
			left := index - prevIndex - 1
			right := i - index - 1
			//fmt.Println(left, right)
			// 计算元素个数 * 最小值 得到res
			res += ((1 + left + right + (right * left)) * arr[index]) % 1000000007
			res = res % 1000000007
		}
		// 满足单调递增栈 直接进栈
		stack = append(stack, i)
	}
	return res
}

// sumSubarrayMinsByDP 暴力遍历 timeout
func sumSubarrayMinsByDP(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sz := len(arr)
	res := 0
	mod := 1000000007
	for i := 0; i < sz; i++ {
		m := arr[i]
		for j := i; j < sz; j++ {
			m = min(m, arr[j])
			res = (res + m) % mod
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
