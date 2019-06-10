package _977

import (
	"fmt"
	"testing"
)

func TestSortedSquares1(t *testing.T) {
	a := []int{-4, -1, 0, 3, 10}
	squares := sortedSquares(a)
	fmt.Println(squares)
}

func TestSortedSquares2(t *testing.T) {
	a := []int{-7, -3, 2, 3, 11}
	squares := sortedSquares(a)
	fmt.Println(squares)
}

func TestSortedSquares3(t *testing.T) {
	a := []int{-2, -1, 3}
	squares := sortedSquares(a)
	fmt.Println(squares)
}
