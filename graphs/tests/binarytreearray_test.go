package tests

import (
	"testing"
	"go-collections/graphs/tree"
)

var tr = tree.NewBTArray()

func TestAddRoot(t *testing.T)  {
	err := tr.AddRoot(17)
	if err != nil {
		t.Error("inserting root returned an error")
	}

	if tr.Size() != 1 {
		t.Error("expected size of tree to be one")
	}

	er := tr.AddRoot(13)
	if er == nil {
		t.Error("Root node is not free. this should not pass")
	}
}

func TestAddLeft(t *testing.T)  {
	err := tr.AddLeft(0, 25)
	if err != nil {
		t.Error("the node specified by 0 is nil")
	}

	if tr.Size() != 2 {
		t.Error("expected size of tree to be 2")
	}
}

func TestAddRight(t *testing.T)  {
	err := tr.AddRight(0, 34)
	if err != nil {
		t.Error("the node specified by 0 is nil")
	}

	if tr.Size() != 3 {
		t.Error("expected size of tree to be 2")
	}

	if tr.Tree()[2] != 34 {
		t.Error("expected right side of root to be 34")
	}
}