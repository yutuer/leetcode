package _432

type AllOne struct {
	maps map[string]*node
	*doubleLinkedList
}

// 更新或者新增操作
type uc interface {
	ucLogic(key string, n *node)
}

// 删除操作
type rc interface {
	rcLogic(key string)
}

func (this *AllOne) ucLogic(key string, n *node) {
	this.maps[key] = n
}

func (this *AllOne) rcLogic(key string) {
	delete(this.maps, key)
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
		this.doubleLinkedList.inc(n, key, this)
	} else {
		this.doubleLinkedList.put(key, this)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	// 查找
	n, exists := this.maps[key]
	if exists {
		this.doubleLinkedList.dec(n, key, this, this)
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
func (this *doubleLinkedList) put(key string, uc uc) {
	// 判断第一个node 是不是 值为1的, 如果不是, 插入一个新生成的node
	next := this.head.next;
	if next.value != 1 {
		n := newNode(1, key)
		this.insertHead(n)

		uc.ucLogic(key, n)
	} else {
		//把n的值放入 node
		next.insertKey(key)

		uc.ucLogic(key, next)
	}

	this.size++
}

// 将n的值加上1
func (this *doubleLinkedList) inc(n *node, key string, uc uc) {
	//一定在map中
	next := n.next

	// 如果后面的节点是+1的, 那么从n中移除key, 放入到后面, 此节点先不删除
	// 否则 判断下是否只有自己, 如果是则只需要换掉值.  如果还有别人, 则生成新节点, 插入n的后面
	if next.value == n.value+1 {
		n.deleteKey(key)
		if len(n.datas) == 0 {
			// 移除掉n
			this.removeNode(n)
		}

		// 插入到下一个节点中
		next.insertKey(key)

		uc.ucLogic(key, next)
	} else {
		if len(n.datas) == 1 {
			// 可以直接将n的value+1
			n.value = n.value + 1
		} else {
			n.deleteKey(key)

			newNode := newNode(n.value+1, key)
			this.insertNode(n, newNode)

			uc.ucLogic(key, newNode)
		}
	}
}

func (this *doubleLinkedList) insertHead(n *node) {
	this.insertNode(this.head, n)
}

func (this *doubleLinkedList) insertNode(prev, n *node) {
	next := prev.next

	// 链接 prev.next 和 n
	n.next = next
	next.prev = n

	// 链接 prev 和 n
	prev.next = n
	n.prev = prev
}

func (this *doubleLinkedList) removeNode(n *node) {
	prev := n.prev
	next := n.next

	prev.next = next
	next.prev = prev
}

// 删除一个值
func (this *doubleLinkedList) remove(key string) {

}

// 将n中的值减去1. 注意此时要和 head 比较
func (this *doubleLinkedList) dec(n *node, key string, rc rc, uc uc) {
	prev := n.prev

	// 如果prev存在, 则插入并更新. 否则插入节点
	if n.value == 1 {
		n.deleteKey(key)

		if len(n.datas) == 0 {
			this.removeNode(n)
		}

		this.size--
	} else {
		if prev.value == n.value-1 {
			prev.insertKey(key)

			n.deleteKey(key)

			if len(n.datas) == 0 {
				this.removeNode(n)
			}

			uc.ucLogic(key, prev)
		} else {
			if len(n.datas) == 1 {
				// 直接将n值-1
				n.value = n.value - 1
			} else {
				newNode := newNode(n.value-1, key)
				this.insertNode(n, newNode)

				uc.ucLogic(key, newNode)
			}
		}
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
