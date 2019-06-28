package _136

import "testing"

func TestA(t *testing.T) {
	result := 1
	duplicate := singleNumber([]int{2, 2, 1})
	if duplicate != result {
		t.Error("need ", result, ", but", duplicate)
	}
}

func TestB(t *testing.T) {
	result := 4
	duplicate := singleNumber([]int{4, 1, 2, 1, 2})
	if duplicate != result {
		t.Error("need ", result, ", but", duplicate)
	}
}
