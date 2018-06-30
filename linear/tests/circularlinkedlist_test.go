package tests

import (
	"go-collections/linear/lists"
	"testing"
)

var cll = new(list.CircularLinkedList)

func TestCllSize(t *testing.T) {
	if cll.Size() != 0 {
		t.Error("Expected size =", 0)
	}
}

func TestCllAddFirst(t *testing.T) {
	cll.AddFirst(5)
	cll.AddFirst(17)
	cll.AddFirst(3)
}

func TestCllFirst(t *testing.T)  {
	if cll.First() != 3 {
		t.Error("Expected first =", 3)
	}
}

func TestCllAddLast(t *testing.T) {
	cll.AddLast(99)
	if cll.Size() != 4 {
		t.Error("Expected size =", 2)
	}
}

func TestCllLast(t *testing.T)  {
	if cll.Last() != 99 {
		t.Error("Expected first =", 99)
	}
}

func TestCllRemoveFirst(t *testing.T)  {
	first := cll.RemoveFirst()
	if first != 3 && cll.First() != 17 {
		t.Error("Expected first =", 17)
	}
}