package _343

import "testing"

func TestA(t *testing.T) {
	n := 2
	r := integerBreak(n)
	if r != 1 {
		t.Error("integerBreak 2, should be ", 1, "but ", r)
	}
}

func TestB(t *testing.T) {
	n := 10
	r := integerBreak(n)
	if r != 36 {
		t.Error("integerBreak 10, should be ", 36, "but ", r)
	}
}
