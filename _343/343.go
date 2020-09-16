package _343

func integerBreak(n int) int {
	// 总乘积
	mul := 1

	// 队列, 防止递归
	q := newQueue()
	q.push(n)

	for ; !q.isEmpty(); {
		next := q.pop()

		if i, b := test(next); b {
			n1, n2 := split(i)

			q.push(n1)
			q.push(n2)
		} else {
			mul *= i
		}
	}
	return mul
}

// 返回值bool表示这个值是否能继续分
func test(n int) (int, bool) {
	if n == 4 {
		return 4, false
	} else if n == 3 {
		return 2, false
	} else if n == 2 {
		return 1, false
	} else if n == 1 {
		return 0, false
	}
	return n, true
}

// 一个int队列
type queue struct {
	datas []int
}

func newQueue() *queue {
	return &queue{datas: make([]int, 0)}
}

func (this *queue) isEmpty() bool {
	return len(this.datas) == 0
}

// 入队
func (this *queue) push(n int) {
	this.datas = append(this.datas, n)
}

// 出队
func (this *queue) pop() int {
	if !this.isEmpty() {
		n := this.datas[0]
		this.datas = this.datas[1:]
		return n
	}
	return 0
}

// 将一个正整数拆分2等分
func split(i int) (n1, n2 int) {
	return i / 2, (i + 2) / 2
}
