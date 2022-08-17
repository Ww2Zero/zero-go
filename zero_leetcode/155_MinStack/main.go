package _55_MinStack

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	dataArr []int
	minArr  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.dataArr = append(this.dataArr, val)
	minSz := len(this.minArr)
	if minSz == 0 || this.minArr[minSz-1] > val {
		this.minArr = append(this.minArr, val)
	} else {
		this.minArr = append(this.minArr, this.minArr[minSz-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.dataArr) == 0 {
		return
	}
	this.dataArr = this.dataArr[:len(this.dataArr)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

func (this *MinStack) Top() int {
	return this.dataArr[len(this.dataArr)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
