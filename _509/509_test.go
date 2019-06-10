package _509

import (
	"testing"
)

func TestFib1(t *testing.T) {
	i := fib(2)
	if i != 1 {
		t.Error("1")
	}
}

func TestFib2(t *testing.T) {
	i := fib(3)
	if i != 2 {
		t.Error("2")
	}
}

func TestFib3(t *testing.T) {
	i := fib(4)
	if i != 3 {
		t.Error("3")
	}
}
