package _217

import "testing"

func TestA(t *testing.T) {
	result := true
	duplicate := containsDuplicate([]int{1, 2, 3, 1})
	if duplicate != result {
		t.Error("need ", result, ", but", duplicate)
	}
}

func TestB(t *testing.T) {
	result := false
	duplicate := containsDuplicate([]int{1, 2, 3, 4})
	if duplicate != result {
		t.Error("need ", result, ", but", duplicate)
	}
}

func TestC(t *testing.T) {
	result := true
	duplicate := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	if duplicate != result {
		t.Error("need ", result, ", but", duplicate)
	}
}
