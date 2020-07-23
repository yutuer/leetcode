package _146

import (
	"testing"
)

func TestLRUCache0(t *testing.T) {
	cache := Constructor(2)

	cache.Put(2, 1)
	cache.Put(2, 2)

	var ret = cache.Get(2)
	if ret != 2 {
		t.Error("answer must equal 2, but:", ret)
	}

	cache.Put(1, 1)
	cache.Put(4, 1)
	ret = cache.Get(2)
	if ret != -1 {
		t.Error("answer must equal -1, but:", ret)
	}
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	var ret = cache.Get(1)
	if ret != 1 {
		t.Error("answer must equal 1, but:", ret)
	}

	cache.Put(3, 3)

	ret = cache.Get(2)
	if ret != -1 {
		t.Error("answer must equal -1, but:", ret)
	}

	cache.Put(4, 4)
	ret = cache.Get(1)
	if ret != -1 {
		t.Error("answer must equal -1, but:", ret)
	}
	ret = cache.Get(3)
	if ret != 3 {
		t.Error("answer must equal 3, but:", ret)
	}
	ret = cache.Get(4)
	if ret != 4 {
		t.Error("answer must equal 4, but:", ret)
	}
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(1)

	cache.Put(2, 1)

	var ret = cache.Get(2)
	if ret != 1 {
		t.Error("answer must equal 1, but:", ret)
	}

	cache.Put(3, 2)

	ret = cache.Get(2)
	if ret != -1 {
		t.Error("answer must equal -1, but:", ret)
	}

	ret = cache.Get(3)
	if ret != 2 {
		t.Error("answer must equal 3, but:", ret)
	}
}

func TestLRUCache3(t *testing.T) {
	cache := Constructor(2)

	cache.Put(2, 1)
	cache.Put(3, 2)

	var ret = cache.Get(3)
	if ret != 2 {
		t.Error("answer must equal 2, but:", ret)
	}
	ret = cache.Get(2)
	if ret != 1 {
		t.Error("answer must equal 1, but:", ret)
	}

	cache.Put(4, 3)

	ret = cache.Get(2)
	if ret != 1 {
		t.Error("answer must equal 1, but:", ret)
	}

	ret = cache.Get(3)
	if ret != -1 {
		t.Error("answer must equal -1, but:", ret)
	}

	ret = cache.Get(4)
	if ret != 3 {
		t.Error("answer must equal 3, but:", ret)
	}
}
