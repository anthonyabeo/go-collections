package tree

type Comparable interface {
	compare(other interface{}) int
}

type node struct {
	item interface{}
	left *node
	right *node
}

func (n *node) compare(other interface{}) int {
	var x int
	switch other.(type) {
		case int:
			if n.item.(int) == other.(int) {
				x = 0
			} else if n.item.(int) > other.(int) {
				x = 1
			} else {
				x = -1
			}
		case string:
			if n.item.(string) == other.(string) {
				x = 0
			} else if n.item.(string) > other.(string) {
				x = 1
			} else {
				x = -1
			}
		case float32:
			if n.item.(float32) == other.(float32) {
				x = 0
			} else if n.item.(float32) > other.(float32) {
				x = 1
			} else {
				x = -1
			}
		case uint:
			if n.item.(uint) == other.(uint) {
				x = 0
			} else if n.item.(uint) > other.(uint) {
				x = 1
			} else {
				x = -1
			}
		case rune:
			if n.item.(rune) == other.(rune) {
				x = 0
			} else if n.item.(rune) > other.(rune) {
				x = 1
			} else {
				x = -1
			}
		case byte:
			if n.item.(byte) == other.(byte) {
				x = 0
			} else if n.item.(byte) > other.(byte) {
				x = 1
			} else {
				x = -1
			}
	}

	return x
}

func (n *node) Item() interface{} {
	return n.item
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) Right() *node {
	return n.right
}


/// Binary Search Tree implementation starts here
/// We define a struct to model the BST. It has a node pointer
/// and number of items for fields. Some methods it implements
/// include Size(), IsEmpty(), Contains() etc
type BinarySearchTree struct {
	root *node
	numItems int
}

func (bst *BinarySearchTree) Root() *node {
	return bst.root
}

func (bst *BinarySearchTree) Size() int {
	return bst.numItems
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.numItems == 0
}

/// Inserts a new item into the BST.
///
/// Args
///		e : item to be inserted
func (bst *BinarySearchTree) Insert(e interface{})  {
	if bst.root == nil {
		bst.root = &node{e, nil, nil}
	} else {
		insert(e, bst.root)
	}
	bst.numItems++
}


func (bst *BinarySearchTree) Remove(e interface{}) interface{} {
	return 0
}

func (bst *BinarySearchTree) Min() interface{} {
	return min(bst.root)
}

func (bst *BinarySearchTree) Max() interface{} {
	return 0
}

func (bst *BinarySearchTree) Contains(e interface{}) bool {
	return false
}

func insert(e interface{}, root *node) {
	if root.compare(e) == 0 {
		return
	} else if root.compare(e) > 0 {
		if root.left == nil {
			root.left = &node{e, nil, nil}
		} else {
			insert(e, root.left)
		}
	} else {
		if root.right == nil {
			root.right = &node{e, nil, nil}
		} else {
			insert(e, root.right)
		}
	}
}

func min(root *node) interface{} {
	if root == nil {
		return nil
	}

	if root.left == nil{
		return root.item
	}

	return min(root.left)
}