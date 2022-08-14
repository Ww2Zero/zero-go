package _46_LRUCache

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}

}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)

}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	// 获取原本head节点的next节点，第一个数据节点
	n := this.head.next
	// 修改next指针的指向逻辑
	// 将当前节点的next指针指向原本第一个数据节点n
	node.next = n
	// 将head节点的next指针指向当前新增节点node
	this.head.next = node
	// 修改prev指针的指向逻辑
	// 将原本第一个数据节点n的prev指针指向当前新增的节点node
	n.prev = node
	// 将当前新增节点node的prev指针指向head节点
	node.prev = this.head
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	//  获取当前节点的上一个节点prevNode
	prevNode := node.prev
	// 修改prevNode的next链接，指向当前节点的下一个节点，打断原本的链接
	prevNode.next = node.next
	// 获取当前节点的下一个节点nextNode
	nextNode := node.next
	// 修改nextNode的prev链接，指向当前节点的上一个节点，打断原本的链接
	nextNode.prev = prevNode
}
