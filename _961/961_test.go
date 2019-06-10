package _961

import "testing"

func Test961_1(t *testing.T) {
	expected := 3

	is := []int{1, 2, 3, 3}

	times := repeatedNTimes(is)
	if times != expected {
		t.Errorf("expected %d, but %d", expected, times)
	}
}

func Test961_2(t *testing.T) {
	expected := 2

	is := []int{2, 1, 2, 5, 3, 2}

	times := repeatedNTimes(is)
	if times != expected {
		t.Errorf("expected %d, but %d", expected, times)
	}
}

func Test961_3(t *testing.T) {
	expected := 5

	is := []int{5, 1, 5, 2, 5, 3, 5, 4}

	times := repeatedNTimes(is)
	if times != expected {
		t.Errorf("expected %d, but %d", expected, times)
	}
}
