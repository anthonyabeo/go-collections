package tests

import (
	"go-collections/graphs/tree"
	"testing"
)

var bst = new(tree.BinarySearchTree)

func TestInsert(t *testing.T)  {
	bst.Insert(5)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)
	if bst.Root().Item() != 5 {
		t.Error("root item should be 5")
	}

	if bst.Root().Left().Item() != 2 {
		t.Error("left value should be 2")
	}

	if bst.Root().Right().Item() != 7 {
		t.Error("Right value should be 7")
	}

	if bst.Size() != 5 {
		t.Error("number of items should be 1")
	}
}

func TestMin(t *testing.T)  {
	min := bst.Min()
	if min != 1 {
		t.Error("min should be 1")
	}
}