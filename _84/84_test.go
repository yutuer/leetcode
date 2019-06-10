package _84

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea1(t *testing.T) {
	a := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(a)
	fmt.Println(area)
}
