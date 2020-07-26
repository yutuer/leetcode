package _09

type CQueue struct {
	s    *stack
	revS *stack
}

func Constructor() CQueue {
	return CQueue{newStack(), newStack()}
}

func (this *CQueue) AppendTail(value int) {
	this.s.push(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.revS.isEmpty() {
		return this.revS.pop()
	}

	if !this.s.isEmpty() {
		for ; !this.s.isEmpty(); {
			this.revS.push(this.s.pop())
		}
	}

	if !this.revS.isEmpty() {
		return this.revS.pop()
	}

	return -1
}

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{make([]int, 0)}
}

func (this *stack) isEmpty() bool {
	return len(this.data) == 0
}

func (this *stack) push(value int) {
	this.data = append(this.data, value)
}

func (this *stack) pop() int {
	if this.isEmpty() {
		return -1
	}

	value := this.data[len(this.data)-1 : len(this.data)]
	this.data = this.data[:len(this.data)-1]
	return value[0]
}
