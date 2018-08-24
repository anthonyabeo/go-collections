package tests

import (
	"go-collections/hashed"
	"testing"
)

var ht = hashed.NewHashTable(13)

func TestNewHT(t *testing.T) {
	if ht.Size() != 0 {
		t.Error("size should be zero")
	}

	if ht.Capacity() != 13 {
		t.Error("capacity is 51")
	}
}

func TestPut(t *testing.T) {
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

func TestGet(t *testing.T) {
	if val, err := ht.Get("b"); err != nil || val == nil {
		t.Error("should be 2")
	}
	if val, err := ht.Get("a"); err != nil || val == nil {
		t.Error("should be 1")
	}
	if val, err := ht.Get("c"); err != nil || val == nil {
		t.Error("should be 3")
	}
	if val, err := ht.Get("d"); err != nil || val == nil {
		t.Error("should be 4")
	}
	if val, err := ht.Get("e"); err != nil || val == nil {
		t.Error("should be 5")
	}
	if val, err := ht.Get("f"); err != nil || val == nil {
		t.Error("should be 6")
	}
	if val, err := ht.Get("g"); err != nil || val == nil {
		t.Error("should be 7")
	}
	if val, err := ht.Get("h"); err != nil || val == nil {
		t.Error("should be 8")
	}
	if val, err := ht.Get("i"); err != nil || val == nil {
		t.Error("should be 9")
	}
	if val, err := ht.Get("j"); err != nil || val == nil {
		t.Error("should be 10")
	}
}

func TestValues(t *testing.T) {
	vals := ht.Values()
	if len(vals) != 10 {
		t.Error("should be 10")
	}

	//for i := 0; i < len(vals); i++ {
	//	if vals[i] == nil {
	//		fmt.Println(nil)
	//	}
	//}
	//fmt.Println(vals)
}

func TestKeySet(t *testing.T) {
	vals := ht.KeySet()
	if len(vals) != 10 {
		t.Error("should be 10")
	}

	//fmt.Println(vals)
}

func TestEntrySet(t *testing.T) {
	vals := ht.EntrySet()
	if len(vals) != 10 {
		t.Error("should be 10")
	}

	//fmt.Println(vals)
}
