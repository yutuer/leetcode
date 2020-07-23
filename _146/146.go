package _146

// 1. 移除的时候别忘了从map中移除
// 2. tail节点更新的时候
// 3. head节点更新

type LRUCache struct {
	maps map[int]*node

	head *node
	tail *node
	size int
	cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*node), newDoubleLinked(capacity)}
}

// 最新的在头部, 最久的在尾部.
type doubleLinked struct {
	head *node
	tail *node
	size int
	cap  int
}

func newDoubleLinked(capacity int) *doubleLinked {
	return &doubleLinked{cap: capacity}
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
func (self *doubleLinked) insertNode(n *node, maps map[int]*node) {
	if self.size == 0 {
		self.head = n
		self.tail = n
		self.head.next = self.tail
		self.tail.prev = self.head

		//更新大小
		self.size++
	} else if self.isFull() {
		//所以当移除的时候,去掉尾部的数据
		// 更新整体指针
		//移除尾部
		// 只有1个满的时候, 直接清空

		delete(maps, self.tail.key)
		self.size--

		if self.size == 1 {
			self.head = nil
			self.tail = nil
		} else {
			prev := self.tail.prev
			prev.next = nil
			self.tail = prev
		}

		self.insertNode(n, maps)
	} else {
		self.insertHead(n)

		//更新大小
		self.size++
	}
}

func (self *doubleLinked) removeTail() {

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
		this.insertNode(newNode, this.maps)

		this.maps[key] = newNode
	} else {
		n.value = value
		this.updateNode(n)
	}
}
