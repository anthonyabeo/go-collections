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
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(4)

	if bst.Root().Item() != 5 {t.Error("root item should be 5")}
	if bst.Root().Left().Item() != 2 {t.Error("left value should be 2")}
	if bst.Root().Right().Item() != 7 {t.Error("Right value should be 7")}
	if bst.Size() != 8 {t.Error("number of items should be 7")}
}

func TestMin(t *testing.T)  {
	min := bst.Min()
	if min != 1 {t.Error("min should be 1")}
}

func TestMax(t *testing.T)  {
	max := bst.Max()
	if max != 9 {t.Error("max should be 9")}
}

func TestContains(t *testing.T)  {
	if _, ok := bst.Contains(9); !ok{
		t.Error("ERROR: tree contains 9")
	}
}

func TestBSTRemove(t *testing.T)  {
	if bst.Remove(2) != 2 {
		t.Error("value not removed")
	}

	if bst.Size() != 7 {
		t.Error("number of items should be 7")
	}

	if bst.Root().Left().Item() != 3 {
		t.Error("new root.left is not 4")
	}
}