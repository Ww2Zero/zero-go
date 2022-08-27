package _25_MyStack

// SqStack single queue stack 。
// 利用队列的先进先出特性，将元素添加到队列的尾部，然后将队列的前面的元素依次添加到队列的尾部，这样就实现了栈的先进后出
type SqStack struct {
	queue []int //创建一个队列
}

// SqStackConstructor  Initialize your data structure here.
func SqStackConstructor() SqStack {
	return SqStack{ //初始化
		queue: make([]int, 0),
	}
}

// Push element x onto stack.
func (this *SqStack) Push(x int) {
	//添加元素，直接添加到队列的尾部
	this.queue = append(this.queue, x)
}

// Pop Removes the element on top of the stack and returns that element.
func (this *SqStack) Pop() int {
	n := len(this.queue) - 1 //判断长度
	for n != 0 {             //除了最后一个，其余的都重新添加到队列里
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		n--
	}
	//弹出元素
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val

}

// Top Get the top element.
func (this *SqStack) Top() int {
	//利用Pop函数，弹出来的元素重新添加
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

// Empty Returns whether the stack is empty.
func (this *SqStack) Empty() bool {
	return len(this.queue) == 0
}
