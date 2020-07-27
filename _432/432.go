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

}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {

}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {

}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {

}

// 双端链表, 用于更改数量的时候使用
type doubleLinkedList struct {
	head, tail *node
	size       int
}

func newDoubleLinkedList() *doubleLinkedList {
	return &doubleLinkedList{&node{-1, nil}, &node{-1, nil}, 0}
}

// 链表节点
type node struct {
	count int             //计数
	datas map[string]byte //拥有的元素
}

func newNode() *node {
	return &node{count: 0, datas: make(map[string]byte)}
}

// 存储一个值
func (this *doubleLinkedList) put(key string) {

}

// 删除一个值
func (this *doubleLinkedList) remove(key string) {

}
