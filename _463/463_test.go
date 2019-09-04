package _463

import "testing"

func TestA(t *testing.T) {
	var a = [][]int{{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0}}
	perimeter := islandPerimeter(a)
	if perimeter != 16 {
		t.Error("a should is 16, but ", perimeter)
	}

	a = [][]int{{1}, {0}}
	perimeter = islandPerimeter(a)
	if perimeter != 4 {
		t.Error("a should is 4, but ", perimeter)
	}

}
