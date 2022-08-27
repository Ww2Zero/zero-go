package _25_MyStack

// DqStack double queue stack 。双队列实现栈
type DqStack struct {
	//创建两个队列
	queue1 []int
	queue2 []int
}

func DqStackConstructor() DqStack {
	return DqStack{ //初始化
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *DqStack) Push(x int) {
	//先将数据存在queue2中
	this.queue2 = append(this.queue2, x)
	//将queue1中所有元素移到queue2中，再将两个队列进行交换
	this.Move()
}

func (this *DqStack) Move() {
	if len(this.queue1) == 0 {
		//交换，queue1置为queue2,queue2置为空
		this.queue1, this.queue2 = this.queue2, this.queue1
	} else {
		//queue1元素从头开始一个一个追加到queue2中
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:] //去除第一个元素
		this.Move()                   //重复
	}
}

func (this *DqStack) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:] //去除第一个元素
	return val

}

func (this *DqStack) Top() int {
	return this.queue1[0] //直接返回
}

func (this *DqStack) Empty() bool {
	return len(this.queue1) == 0
}
