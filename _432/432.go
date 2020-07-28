package _432

type AllOne struct {
	maps map[string]*node
	*doubleLinkedList
}

type uc interface {
	ucLogic(key string, n *node)
}

func (this *AllOne) ucLogic(key string, n *node) {
	this.maps[key] = n
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{maps: make(map[string]*node), doubleLinkedList: newDoubleLinkedList()}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	// 查找
	n, exists := this.maps[key]
	if exists {
		this.doubleLinkedList.add(n, key, this)
	} else {
		this.doubleLinkedList.put(key)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	// 查找
	n, exists := this.maps[key]
	if exists {
		this.doubleLinkedList.sub(n, key)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if (this.doubleLinkedList.size == 0) {
		return ""
	}

	prev := this.tail.prev
	for k, _ := range prev.datas {
		return k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if (this.doubleLinkedList.size == 0) {
		return ""
	}

	next := this.head.next
	for k, _ := range next.datas {
		return k
	}
	return ""
}

// 双端链表, 用于更改数量的时候使用
type doubleLinkedList struct {
	head, tail *node
	size       int
}

func newDoubleLinkedList() *doubleLinkedList {
	head, tail := newNode(0, ""), newNode(0, "")
	head.next = tail
	tail.prev = head
	return &doubleLinkedList{head, tail, 0}
}

// 存储一个值, 放到第一个位置
func (this *doubleLinkedList) put(key string) {
	// 判断第一个node 是不是 值为1的, 如果不是, 插入一个新生成的node
	next := this.head.next;
	if next.value != 1 {
		n := newNode(1, key)
		this.insertHead(n)
	} else {
		//把n的值放入 node
		next.insertKey(key)
	}
}

func (this *doubleLinkedList) insertHead(n *node) {
	this.insertNode(this.head, n)
}

func (this *doubleLinkedList) insertNode(prev, n *node) {
	next := prev.next

	// 链接 head.next 和 n
	n.next = next
	next.prev = n

	// 链接head和n

	this.head.next = n
	n.prev = this.head
}

// 删除一个值
func (this *doubleLinkedList) remove(key string) {

}

// 将n中的值减去1
func (this *doubleLinkedList) sub(n *node, key string) {

}

// 将n的值加上1
func (this *doubleLinkedList) add(n *node, key string, uc uc) {
	//一定在map中
	next := n.next

	// 如果后面的节点是+1的, 那么从n中移除key, 放入到后面, 此节点先不删除
	// 否则 判断下是否只有自己, 如果是则只需要换掉值.  如果还有别人, 则生成新节点, 插入n的后面
	n.deleteKey(key)
	if next.value == n.value+1 {
		next.insertKey(key)
	} else {
		newNode := newNode(n.value+1, key)
		this.insertNode(n, newNode)
	}

}

// 链表节点
type node struct {
	value      int             //计数
	datas      map[string]byte //拥有的元素
	prev, next *node
}

func newNode(value int, key string) *node {
	maps := make(map[string]byte)
	maps[key] = '1'

	return &node{value: value, datas: maps, prev: nil, next: nil}
}

func (this *node) insertKey(key string) {
	this.datas[key] = '1'
}
func (this *node) deleteKey(key string) {
	delete(this.datas, key)
}
