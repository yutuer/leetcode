package _206

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	var input = transfer([]int{1, 2, 3, 4, 5})

	fmt.Println(input)

	output := reverseList(input)
	fmt.Println(output)
}
