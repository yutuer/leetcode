package _832

import (
	"fmt"
	"testing"
)

func TestFlipAndInvertImage1(t *testing.T) {
	a := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}

	flipAndInvertImage(a)

	fmt.Println(a)
}

func TestFlipAndInvertImage2(t *testing.T) {
	a := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}

	flipAndInvertImage(a)

	fmt.Println(a)
}

func TestFlipAndInvertImage3(t *testing.T) {
	a := [][]int{{1}}

	flipAndInvertImage(a)

	fmt.Println(a)
}

