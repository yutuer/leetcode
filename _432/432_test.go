package _432

import "testing"

func TestA(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("hello")

	maxKey := allOne.GetMaxKey()
	if maxKey != "hello" {
		t.Errorf("expect hello, but [%s]", maxKey)
	}

	minKey := allOne.GetMinKey()
	if minKey != "hello" {
		t.Errorf("expect hello, but [%s]", minKey)
	}

}
func TestB(t *testing.T) {

	allOne := Constructor()
	allOne.Inc("hello")

	allOne.Inc("hello")
	allOne.Inc("hello")

	maxKey := allOne.GetMaxKey()
	if maxKey != "hello" {
		t.Errorf("expect hello, but [%s]", maxKey)
	}

	minKey := allOne.GetMinKey()
	if minKey != "hello" {
		t.Errorf("expect hello, but [%s]", minKey)
	}

	allOne.Inc("leet")

	maxKey = allOne.GetMaxKey()
	if maxKey != "hello" {
		t.Errorf("expect hello, but [%s]", maxKey)
	}

	minKey = allOne.GetMinKey()
	if minKey != "leet" {
		t.Errorf("expect leet, but [%s]", minKey)
	}
}
