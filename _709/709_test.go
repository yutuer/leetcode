package _709

import "testing"

func Test709_1(t *testing.T) {
	s := "Hello"

	rs := toLowerCase(s)
	if rs != "hello" {
		t.Errorf("expect:%s but:%s", "hello", rs)
	}
}

func Test709_2(t *testing.T) {
	s := "here"

	rs := toLowerCase(s)
	if rs != "here" {
		t.Errorf("expect:%s but:%s", "here", rs)
	}
}

func Test709_3(t *testing.T) {
	s := "LOVELY"

	rs := toLowerCase(s)
	if rs != "lovely" {
		t.Errorf("expect:%s but:%s", "lovely", rs)
	}
}

func Test709_4(t *testing.T) {
	s := "al&phaBET"

	rs := toLowerCase(s)
	if rs != "al&phabet" {
		t.Errorf("expect:%s but:%s", "al&phabet", rs)
	}
}
