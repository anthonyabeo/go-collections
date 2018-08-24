package linked

import "sync"

/// This is Node implementation. The node is private
/// outside this file since it starts with lowercase
/// Here a struct data structure to hold the items
/// of the Node. These include the item and the next
type scalingNode struct {
	item interface{}
	next *scalingNode
	lock *sync.Mutex
}

// returns a pointer to the item in the node
func (n *scalingNode) SItem() interface{} {
	return n.item
}

// returns a pointer to the next node in the list
func (n *scalingNode) SNext() *scalingNode {
	return n.next
}

/// This is SinglyLikedList implementation
/// Here a struct data structure to hold the items
/// of the list. These include the number of items,
/// the head and the tail of the list
type ScalingSinglyLinkedList struct {
	numItems int
	head     *scalingNode
	tail     *scalingNode
}

// Add a new item to the front of the list
func (sll *ScalingSinglyLinkedList) SAddFirst(e interface{}) {
	sll.head = &scalingNode{item: e, next: sll.head, lock: new(sync.Mutex)}
	if sll.numItems == 0 {
		sll.tail = sll.head
	}
	sll.numItems += 1
}

// returns the first element from the list
func (sll *ScalingSinglyLinkedList) SFirst() interface{} {
	if sll.SIsEmpty() {
		return nil
	}

	sll.head.lock.Lock()
	defer sll.head.lock.Unlock()

	return sll.head.item
}

// Add a new item to the end of the list
func (sll *ScalingSinglyLinkedList) SAddLast(e interface{}) {

	newNode := &scalingNode{item: e, next: sll.head, lock: new(sync.Mutex)}
	if sll.SIsEmpty() {
		sll.head = newNode
	} else {
		sll.tail.next = newNode
	}

	sll.tail = newNode
	sll.numItems += 1
}

// returns the first element from the list
func (sll *ScalingSinglyLinkedList) SLast() interface{} {

	if sll.SIsEmpty() {
		return nil
	}

	sll.tail.lock.Lock()
	defer sll.tail.lock.Unlock()

	return sll.tail.item
}

// adds a new item after the specified node
func (sll *ScalingSinglyLinkedList) SAddAfter(n *scalingNode, e interface{}) {

	if n != nil {
		n.lock.Lock()

		n.next.lock.Lock()
		defer n.lock.Unlock()

		n.next = &scalingNode{item: e, next: sll.head, lock: new(sync.Mutex)}
		sll.numItems++

		n.lock.Unlock()
	}
}

// Returns the number of items in the list
func (sll *ScalingSinglyLinkedList) SSize() int {

	return sll.numItems
}

// returns a boolean indicating whether or not
// the list is empty or not
func (sll *ScalingSinglyLinkedList) SIsEmpty() bool {

	return sll.numItems == 0
}

// removes and returns the first item in the list
func (sll *ScalingSinglyLinkedList) SRemoveFirst() interface{} {
	sll.head.lock.Lock()
	defer sll.head.lock.Unlock()

	if sll.SIsEmpty() {
		return nil
	}
	first := sll.SFirst()
	sll.head = sll.head.SNext()
	sll.numItems--

	if sll.numItems == 0 {
		sll.tail = nil
	}

	return first
}
