package tests

import (
	// Built-in packages
	"fmt"
	"testing"

	// User defined packages
	"go-collections/linear/lists"
)

var sll = new(list.SinglyLinkedList)

func TestSize(t *testing.T) {
	if sll.Size() != 0 {
		t.Error("Expected size =", 0)
	}
}

func TestAddFirst(t *testing.T) {
	sll.AddFirst(5)
	sll.AddFirst(17)
	sll.AddFirst(3)
	fmt.Println("Number of items: ", sll.Size())
}

func TestFirst(t *testing.T) {
	if sll.First() != 3 {
		t.Error("Expected first =", 3)
	}
}

func TestAddLast(t *testing.T) {
	sll.AddLast(99)
	fmt.Println("Number of items: ", sll.Size())
}

func TestLast(t *testing.T) {
	if sll.Last() != 99 {
		t.Error("Expected first =", 99)
	}
}

func TestRemoveFirst(t *testing.T) {
	first := sll.RemoveFirst()
	if first != 3 && sll.First() != 17 {
		t.Error("Expected first =", 17)
	}
}
