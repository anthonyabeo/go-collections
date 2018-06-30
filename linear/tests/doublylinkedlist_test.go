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

func TestDllAddLast(t *testing.T) {
	dll.AddLast("knight")
	dll.AddLast("pawn")
	dll.AddLast("rookie")

	if dll.Size() != 5 {
		t.Error("Expected size of dll to be", 5)
	}

	if dll.Last() != "rookie" {
		t.Error("Expected first item in dll to be queen")
	}
}

func TestDllRemoveFirst(t *testing.T)  {
	item := dll.RemoveFirst()

	if item != "queen" {
		t.Error("Expected queen to be remove")
	}

	if dll.Size() != 4 {
		t.Error("Expected the number of item to be 4")
	}
}

func TestDllRemoveLast(t *testing.T)  {
	item := dll.RemoveLast()
	dll.RemoveLast()
	dll.RemoveLast()

	if item != "rookie" {
		t.Error("Expected rookie to be remove")
	}

	if dll.Size() != 1 {
		t.Error("Expected the number of item to be 3")
	}
}