package tests

import (
	"go-collections/hashed"
	"testing"
)

var hs = hashed.NewHashSet(13)

func TestAdd(t *testing.T) {
	if !hs.IsEmpty() {
		t.Error("should be empty")
	}

	hs.Add("a")
	hs.Add("b")
	hs.Add("c")
	hs.Add("d")
	hs.Add("e")
	hs.Add("f")
	hs.Add("g")
	hs.Add("h")
	hs.Add("i")
	hs.Add("j")

	if hs.Size() != 10 {
		t.Error("size should be 10")
	}
}

func TestContains(t *testing.T) {
	if !hs.Contains("e") {
		t.Error("should contain e")
	}

	if hs.Contains(4) {
		t.Error("should NOT contain 4")
	}
}

func TestRemove(t *testing.T) {
	if val, err := hs.Remove("b"); err != nil || val == nil {
		t.Error("should be 2")
	}
	if val, err := hs.Remove("a"); err != nil || val == nil {
		t.Error("should be 1")
	}
	if val, err := hs.Remove("c"); err != nil || val == nil {
		t.Error("should be 3")
	}
	if val, err := hs.Remove("d"); err != nil || val == nil {
		t.Error("should be 4")
	}

	if hs.Size() != 6 {
		t.Error("size should be 6")
	}
}

func TestEmpty(t *testing.T) {
	hs.Empty()
	if hs.Size() != 0 {
		t.Error("size should be zero")
	}

	if !hs.IsEmpty() {
		t.Error("capacity is 51")
	}
}
