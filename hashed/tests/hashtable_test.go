package tests

import (
	"go-collections/hashed"
	"testing"
)

var ht = hashed.NewHashTable()

func TestNewHT(t *testing.T)  {
	if ht.Size() != 0 {
		t.Error("size should be zero")
	}

	if ht.Capacity() != 13 {
		t.Error("capacity is 51")
	}
}

func TestPut(t *testing.T)  {
	ht.Put("a", 1)
	ht.Put("b", 2)
	ht.Put("c", 3)
	ht.Put("d", 4)
	ht.Put("e", 5)
	ht.Put("f", 6)
	ht.Put("g", 7)
	ht.Put("h", 8)
	ht.Put("i", 9)
	ht.Put("j", 10)

	if ht.Size() != 10 {
		t.Error("size should be 10")
	}

	if ht.Capacity() != 26 {
		t.Error("capacity should be 26")
	}
}

func TestGet(t *testing.T)  {
	if ht.Get("b") != 2 {
		t.Error("should be 2")
	}

	if ht.Get("a") != 1 {
		t.Error("should be 2")
	}

	if ht.Get("c") != 3 {
		t.Error("should be 2")
	}
}