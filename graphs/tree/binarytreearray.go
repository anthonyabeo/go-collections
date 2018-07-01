package tree

import (
	"errors"
)

type BinaryTreeArray struct {
	tree []interface{}
	numItems int
}

func NewBTArray() *BinaryTreeArray {
	return &BinaryTreeArray{
		make([]interface{}, 10),
		0,
	}
}

// Returns the number of items in the tree
func (tree *BinaryTreeArray) Size() int {
	return tree.numItems
}

// Returns a boolean indicating whether or not
// the tree has no items in it.
func (tree *BinaryTreeArray) IsEmpty() bool {
	return tree.numItems == 0
}

/// Inserts a new item into the tree.
/// The item becomes the root node of the tree. An error
/// occurs if the tree is not empty
///
/// Args
/// e interface{} : the item to be inserted
///
/// Returns
/// error : nil if the tree is empty
///       : error message if a root node already exists
func (tree *BinaryTreeArray) AddRoot(e interface{}) (error) {
	if !tree.IsEmpty() && tree.tree[0] != nil {
		return errors.New("root node already exist")
	}
	tree.tree[0] = e
	tree.numItems++

	return nil
}

/// Inserts a new item into the tree.
/// The item becomes the left child of the element, p
///
/// Args
/// 	p int : the index int the array of the parent node
/// 	e interface{} : the item to be inserted
///
/// Returns
/// 	error: if the index p points to nil
func (tree *BinaryTreeArray) AddLeft(p int, e interface{}) error {
	 if tree.tree[p] == nil {
	 	return errors.New("cannot assign left value to nil")
	 }

	 leftIndex := (2 * p) + 1
	 tree.tree[leftIndex] = e
	 tree.numItems++

	 return nil
}

func (tree *BinaryTreeArray) Tree() []interface{} {
	return tree.tree
}

/// Inserts a new item into the tree.
/// The item becomes the right child of the element, p
///
/// Args
/// 	p int : the index int the array of the parent node
/// 	e interface{} : the item to be inserted
func (tree *BinaryTreeArray) AddRight(p int, e interface{}) error {
	if tree.tree[p] == nil {
		return errors.New("cannot assign left value to nil")
	}

	leftIndex := (2 * p) + 2
	tree.tree[leftIndex] = e
	tree.numItems++

	return nil
}

/// Performs BFS or DFS on the tree and returns index of the
/// item specified.
///
/// Args
/// 	e interface{} : the item to be inserted
///
/// Returns
/// 	int: index of the element in the tree
///		error: if the value does not exist in the tree
func (tree *BinaryTreeArray) Get(e interface{}) (int, error) {
	for i := 0; i < tree.Size(); i++  {
		if tree.tree[i] == e {
			return i, nil
		}
	}

	return -1, errors.New("no such item")
}

/// Updates the value at index p with e.
/// It returns the previously stored element
///
/// Args
/// 	p int : the index int the array of the parent node
/// 	e interface{} : the item to be inserted
///
/// Returns
///		error: if the value at p does not exist in the tree
func (tree *BinaryTreeArray) Set(p int, e interface{}) error {
	if tree.tree[p] == nil {
		return errors.New("cannot assign left value to nil")
	}

	tree.tree[p] = e

	return nil
}

/// Removes the node at index p, replacing it with its
/// child (if any) and returns the element that was stored at p
/// if p has two children, return an error.
///
/// Args
/// 	p int : the index in the array of the value to be removed
func (tree *BinaryTreeArray) Remove(p int) error {
	if tree.tree[p] == nil {
		return errors.New("cannot assign left value to nil")
	}

	if tree.tree[(2*p) + 1] != nil && tree.tree[(2*p) + 2] != nil {

		return errors.New("the value has two children")

	} else if tree.tree[(2*p) + 1] != nil && tree.tree[(2*p) + 2] == nil {
		tree.tree[p] = tree.tree[(2*p) + 1]
		tree.tree[(2*p) + 1] = nil
	} else if tree.tree[(2*p) + 1] == nil && tree.tree[(2*p) + 2] != nil {
		tree.tree[p] = tree.tree[(2*p) + 2]
		tree.tree[(2*p) + 2] = nil
	} else {
		tree.tree[p] = nil
	}

	tree.numItems--

	return nil
}