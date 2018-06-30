package tests

import (
	"testing"
	"go-collections/linear"
)

var deque = linear.NewDeque()

func TestDequeSize(t *testing.T) {
	if deque.Size() != 0 {
		t.Error("Expected size =", 0)
	}
}

func TestDequeIsEmpty(t *testing.T) {
	if !deque.IsEmpty() {
		t.Error("list is not empty, but it should")
	}
}

func TestDequeAddFirst(t *testing.T) {
	deque.AddFirst("king")
	deque.AddFirst("queen")

	if deque.Size() != 2 {
		t.Error("Expected size of dll to be", 2)
	}

	if deque.First() != "queen" {
		t.Error("Expected first item in dll to be queen")
	}

	if deque.Last() != "king" {
		t.Error("Expected last item in dll to be king")
	}
}

func TestDequeAddLast(t *testing.T) {
	deque.AddLast("knight")
	deque.AddLast("pawn")
	deque.AddLast("rookie")

	if deque.Size() != 5 {
		t.Error("Expected size of deque to be", 5)
	}

	if deque.Last() != "rookie" {
		t.Error("Expected first item in deque to be queen")
	}
}

func TestDequeRemoveFirst(t *testing.T)  {
	item := deque.RemoveFirst()

	if item != "queen" {
		t.Error("Expected queen to be remove")
	}

	if deque.Size() != 4 {
		t.Error("Expected the number of item to be 4")
	}
}

func TestDequeRemoveLast(t *testing.T)  {
	item := deque.RemoveLast()
	deque.RemoveLast()
	deque.RemoveLast()

	if item != "rookie" {
		t.Error("Expected rookie to be remove")
	}

	if deque.Size() != 1 {
		t.Error("Expected the number of item to be 3")
	}
}