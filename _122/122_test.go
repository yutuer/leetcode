package _122

import "testing"

func TestA(t *testing.T) {
	result := 7
	n := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if n != result {
		t.Errorf("right:%d, you:%d", result, n)
	}
}

func TestB(t *testing.T) {
	result := 4
	n := maxProfit([]int{1, 2, 3, 4, 5})
	if n != result {
		t.Errorf("right:%d, you:%d", result, n)
	}
}

func TestC(t *testing.T) {
	result := 0
	n := maxProfit([]int{7, 6, 4, 3, 1})
	if n != result {
		t.Errorf("right:%d, you:%d", result, n)
	}
}
