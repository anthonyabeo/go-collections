package tests

import (
	"go-collections/linear/lists"
	"testing"
)

var dll = list.NewDll()

func TestDllSize(t *testing.T) {
	if dll.Size() != 0 {
		t.Error("Expected size =", 0)
	}
}

func TestDllIsEmpty(t *testing.T) {
	if !dll.IsEmpty() {
		t.Error("list is not empty, but it should")
	}
}

func TestDllAddFirst(t *testing.T) {
	dll.AddFirst("king")
	dll.AddFirst("queen")

	if dll.Size() != 2 {
		t.Error("Expected size of dll to be", 2)
	}

	if dll.First() != "queen" {
		t.Error("Expected first item in dll to be queen")
	}

	if dll.Last() != "king" {
		t.Error("Expected last item in dll to be king")
	}
}