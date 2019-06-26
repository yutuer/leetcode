package _476

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	complement := findComplement(5)
	fmt.Println(complement)
	complement = findComplement(1)
	fmt.Println(complement)
	complement = findComplement(0)
	fmt.Println(complement)
}