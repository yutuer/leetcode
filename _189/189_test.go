package _189

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println(arr)
}

func TestB(t *testing.T) {
	arr := []int{2147483647, -2147483648, 33, 219, 0}
	rotate(arr, 4)
	fmt.Println(arr)
}
