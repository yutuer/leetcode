package _461

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	x, y := 1, 4
	distance := hammingDistance(x, y)
	fmt.Println(distance)
}
