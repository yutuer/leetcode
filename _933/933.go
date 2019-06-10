package _933

import "sort"

type RecentCounter struct {
	ping  []int
	index int
}

func Constructor() RecentCounter {
	return RecentCounter{ping: make([]int, 10000), index: 0}
}

var total int

func (this *RecentCounter) Ping(t int) int {
	this.ping[this.index] = t
	this.index ++

	total = 0
	for i := len(this.ping) - 1; i >= 0; i-- {
		if t-this.ping[i] > 3000 {
			break
		}
		total ++
	}

	sort.SearchInts(this.ping,t-3000)

	return total
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
