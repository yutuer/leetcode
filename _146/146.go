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

type oc interface {
	doOcLogic(n *node)
}

type ic interface {
	doIcLogic(n *node)
}

func (self *LRUCache) doOcLogic(n *node) {
	delete(self.maps, n.key)
}

func (self *LRUCache) doIcLogic(n *node) {
	self.maps[n.key] = n
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
func (self *doubleLinked) updateNode(n *node, oc oc, ic ic) {
	if n != nil {
		self.removeNode(n, oc)
		self.insertHead(n, oc, ic)
	}
}

// 插入一个节点
func (self *doubleLinked) insertNode(prev, n *node, oc oc, ic ic) {
	if n == nil || prev == nil {
		return
	}

	if self.isFull() {
		self.removeTail(oc)
	}

	next := prev.next

	// 链接n和next
	n.next = next
	next.prev = n

	// 链接n和prev
	prev.next = n
	n.prev = prev

	self.size++

	ic.doIcLogic(n)
}

// 插入一个node到头部
func (self *doubleLinked) insertHead(n *node, oc oc, ic ic) {
	self.insertNode(self.head, n, oc, ic)
}

// 移除一个节点
func (self *doubleLinked) removeNode(n *node, ic oc) {
	if self.size == 0 {
		return
	}

	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev

	self.size--

	ic.doOcLogic(n)
}

// 从尾部移除一个节点
func (self *doubleLinked) removeTail(ic oc) {
	self.removeNode(self.tail.prev, ic)
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
		this.updateNode(n, this, this)
		return n.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	n := this.maps[key]
	if n == nil {
		newNode := newNode(key, value)
		this.insertHead(newNode, this, this)
		this.maps[key] = newNode
	} else {
		n.value = value
		this.updateNode(n, this, this)
	}
}
