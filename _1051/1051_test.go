package _1051

import "testing"

func TestA(t *testing.T) {
	total := heightChecker([]int{1, 1, 4, 2, 1, 3})
	if total != 3 {
		t.Error("total should 3, but", total)
	}
}
