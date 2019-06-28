package _26

import "testing"

func TestA(t *testing.T) {
	n := removeDuplicates([]int{1, 1, 2})
	if n != 2 {
		t.Error("n should 2, but ", n)
	}
}

func TestB(t *testing.T) {
	n := removeDuplicates([]int{1, 1})
	if n != 1 {
		t.Error("n should 1, but ", n)
	}
}

func TestC(t *testing.T) {
	n := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if n != 5 {
		t.Error("n should 5, but ", n)
	}
}
