package _852

import (
	"fmt"
	"testing"
)

func TestPeakIndexInMountainArray1(t *testing.T) {
	array := peakIndexInMountainArray([]int{0, 1, 0})
	fmt.Println(array)
}

func TestPeakIndexInMountainArray2(t *testing.T) {
	array := peakIndexInMountainArray([]int{0, 2, 1, 0})
	fmt.Println(array)
}
