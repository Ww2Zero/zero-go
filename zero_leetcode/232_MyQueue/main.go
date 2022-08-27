package _32_MyQueue

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

//Constructor  Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

// Push element x to the stackOut of queue.
func (this *MyQueue) Push(x int) {
	for len(this.stackOut) != 0 {
		val := this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
		this.stackIn = append(this.stackIn, val)
	}
	this.stackIn = append(this.stackIn, x)
}

// Pop  Removes the element from in front of queue and returns that element.
func (this *MyQueue) Pop() int {
	for len(this.stackIn) != 0 {
		val := this.stackIn[len(this.stackIn)-1]
		this.stackIn = this.stackIn[:len(this.stackIn)-1]
		this.stackOut = append(this.stackOut, val)
	}
	if len(this.stackOut) == 0 {
		return 0
	}
	val := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return val
}

// Peek  Get the front element.
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == 0 {
		return 0
	}
	this.stackOut = append(this.stackOut, val)
	return val
}

// Empty  Returns whether the queue is empty.
func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
