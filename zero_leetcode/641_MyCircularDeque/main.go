package _41_MyCircularDeque

//leetcode submit region begin(Prohibit modification and deletion)

//设计实现双端队列。
//
//实现 MyCircularDeque 类:
//
//MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
//boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
//boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
//boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
//boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
//int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
//int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
//boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false 。
//boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type MyCircularDeque struct {
	capacity int
	size     int
	head     *Node
	tail     *Node
}

func Constructor(k int) MyCircularDeque {
	deque := MyCircularDeque{
		capacity: k,
		size:     0,
		head:     &Node{value: -1},
		tail:     &Node{value: -1},
	}
	deque.head.next = deque.tail
	deque.tail.prev = deque.head
	return deque
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size == this.capacity {
		return false
	}
	node := &Node{value: value}

	next := this.head.next

	node.next = next
	this.head.next = node

	next.prev = node
	node.prev = this.head
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size == this.capacity {
		return false
	}
	node := &Node{value: value}

	prev := this.tail.prev

	prev.next = node
	node.next = this.tail

	this.tail.prev = node
	node.prev = prev
	this.size++
	return true

}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	front := this.head.next

	this.head.next = front.next
	front.next.prev = this.head

	front.next = nil
	front.prev = nil

	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	prev := this.tail.prev

	this.tail.prev = prev.prev
	prev.prev.next = this.tail

	prev.next = nil
	prev.prev = nil

	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	front := this.head.next
	return front.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	prev := this.tail.prev
	return prev.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
