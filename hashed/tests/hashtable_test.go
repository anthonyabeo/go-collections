package tests

import (
	"go-collections/hashed"
	"testing"
	"fmt"
)

var ht = hashed.NewHashTable(13)

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
	if ht.Get("b") != 2 {t.Error("should be 2")}
	if ht.Get("a") != 1 {t.Error("should be 1")}
	if ht.Get("c") != 3 {t.Error("should be 3")}
	if ht.Get("d") != 4 {t.Error("should be 4")}
	if ht.Get("e") != 5 {t.Error("should be 5")}
	if ht.Get("f") != 6 {t.Error("should be 6")}
	if ht.Get("g") != 7 {t.Error("should be 7")}
	if ht.Get("h") != 8 {t.Error("should be 8")}
	if ht.Get("i") != 9 {t.Error("should be 9")}
	if ht.Get("j") != 10 {t.Error("should be 10")}
}

func TestValues(t *testing.T)  {
	vals := ht.Values()
	if len(vals) != 10 {
		t.Error("should be 10")
	}

	for i := 0; i < len(vals); i++ {
		if vals[i] == nil {
			fmt.Println(nil)
		}
	}
	fmt.Println(vals)
}

func TestKeySet(t *testing.T)  {
	vals := ht.KeySet()
	if len(vals) != 10 {
		t.Error("should be 10")
	}

	fmt.Println(vals)
}

func TestEntrySet(t *testing.T)  {
	vals := ht.EntrySet()
	if len(vals) != 10 {
		t.Error("should be 10")
	}

	fmt.Println(vals)
}

