package _1051

import (
	"sort"
)

func heightChecker(heights []int) int {
	var h = make([]int, len(heights))
	_ = copy(h, heights)
	sort.Ints(heights)

	var total int
	for i := 0; i < len(h); i++ {
		if heights[i] != h[i] {
			total++
		}
	}
	return total
}
