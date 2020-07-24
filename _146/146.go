package _146

// 1. 移除的时候别忘了从map中移除

type LRUCache struct {
	maps map[int]*node
	// 双向链表
	*doubleLinked
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*node), newDoubleLinked(capacity)}
}

// 构造Fake的头尾节点
// 最新的在头部, 最久的在尾部.
type doubleLinked struct {
	head *node
	tail *node
	size int
	cap  int
}

func newDoubleLinked(capacity int) *doubleLinked {
	//开始构造2个, 好像就不用判断长度了
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return &doubleLinked{head: head, tail: tail, size: 0, cap: capacity}
}

func (self *doubleLinked) isFull() bool {
	return self.size == self.cap
}

// 更新node
func (self *doubleLinked) updateNode(n *node) {
	if self.size > 1 {
		prev := n.prev
		next := n.next
		if prev != nil {
			prev.next = next
		}

		if next != nil {
			next.prev = prev
		}

		self.insertHead(n)

		if n == self.tail {
			self.tail = prev
		}
	}
}

// 插入node
func (self *doubleLinked) insertNode(n *node) {
	head := self.head
	h := head.next

	n.next = h
	h.prev = n

	// 链接新头
	head.next = n
	n.prev = head

	//更新大小
	self.size++
}

// 移除尾部
func (self *doubleLinked) removeNode(n *node) bool {
	if (self.size == 0) {
		return false
	}

	return true
}

func (self *doubleLinked) insertHead(n *node) {
	n.prev = nil
	n.next = self.head
	self.head.prev = n
	self.head = n
}

type node struct {
	prev  *node
	next  *node
	key   int
	value int
}

func newNode(key, v int) *node {
	return &node{key: key, value: v}
}

func (this *LRUCache) Get(key int) int {
	n := this.maps[key]
	if n == nil {
		return -1
	} else {
		this.updateNode(n)
		return n.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	n := this.maps[key]
	if n == nil {
		newNode := newNode(key, value)
		this.insertNode(newNode)

		this.maps[key] = newNode
	} else {
		n.value = value
		this.updateNode(n)
	}
}
