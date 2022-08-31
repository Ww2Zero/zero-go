package _39_dailyTemperatures

// dailyTemperatures 每一天的温度
//给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，
//其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。
//如果气温在这之后都不会升高，请在该位置用0 来
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	// 初始化栈顶的第一个下表的索引0，初始化栈顶的索引数据为索引0.
	// stack中记录的是不同天气气温单调增高的的天数的索引
	stack := []int{0}

	for i := 0; i < len(temperatures); i++ {
		// 栈顶的索引数据top
		top := stack[len(stack)-1]

		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			// 当前的温度(temperatures[i]) > 栈顶的温度(temperatures[top])，则记录
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				// 记录top这个索引的位置下，下一个比当前位置top的温度更高的索引的位置距离当前位置的距离。（i-top）:目标
				res[top] = i - top
				// 移除栈顶数据
				stack = stack[:len(stack)-1]
				// 更新栈顶top的值
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			// 数据都比较移除完成之后，将当前的索引i加入栈
			stack = append(stack, i)
		}
	}
	return res
}
