package _432

type AllOne struct {
	maps map[string]*node
	*doubleLinkedList
}

type oc interface {
	ocLogic()
}

type ic interface {
	icLogic()
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
		this.doubleLinkedList.add(n)
	} else {
		this.doubleLinkedList.put(key)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	// 查找
	n, exists := this.maps[key]
	if exists {
		this.doubleLinkedList.sub(n)
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
	head, tail := newNode(), newNode()
	head.next = tail
	tail.prev = head
	return &doubleLinkedList{head, tail, 0}
}

// 链表节点
type node struct {
	count      int             //计数
	datas      map[string]byte //拥有的元素
	prev, next *node
}

func newNode() *node {
	return &node{count: 0, datas: make(map[string]byte), prev: nil, next: nil}
}

// 存储一个值
func (this *doubleLinkedList) put(key string) {

}

// 删除一个值
func (this *doubleLinkedList) remove(key string) {

}

func (this *doubleLinkedList) sub(n *node) {

}

func (this *doubleLinkedList) add(n *node) {

}
