package list

// usual imports
import "fmt"


/// This is Node implementation. The node is private
/// outside this file since it starts with lowercase
/// Here a struct data structure to hold the items
/// of the Node. These include the item and the next
type node struct {
    item interface{}
    next *node
}

// returns a pointer to the item in the node
func (n node) Item() interface{} {
    return n.item
}

// returns a pointer to the next node in the list
func (n node) Next() *node {
    return n.next
}


/// This is SinglyLikedList implementation
/// Here a struct data structure to hold the items
/// of the list. These include the number of items,
/// the head and the tail of the list
type SinglyLinkedList struct {
    numItems int
    head *node
    tail *node
}

// Add a new item to the front of the list
func (sll *SinglyLinkedList) AddFirst(e interface{}) {
    sll.head = &node{item: e, next: sll.head}
    if sll.numItems == 0 {
        sll.tail = sll.head
    }
    sll.numItems += 1
}

// returns the first element from the list
func (sll *SinglyLinkedList) First() interface{}  {
	if sll.IsEmpty() {
		return nil
	}
    return sll.head.item
}

// Add a new item to the end of the list
func (sll *SinglyLinkedList) AddLast(e interface{}) {
    newNode := &node{item: e, next: nil}
    if sll.IsEmpty() {
        sll.head = newNode
    } else {
        sll.tail.next = newNode
    }

    sll.tail = newNode
    sll.numItems += 1
}

// returns the first element from the list
func (sll *SinglyLinkedList) Last() interface{} {
	if sll.IsEmpty() {
		return nil
	}
    return sll.tail.item
}

// Returns the number of items in the list
func (sll *SinglyLinkedList) Size() int {
    return sll.numItems
}

// returns a boolean indicating whether or not
// the list is empty or not
func (sll *SinglyLinkedList) IsEmpty() bool  {
    return sll.numItems == 0
}

// removes and returns the first item in the list
func (sll *SinglyLinkedList) RemoveFirst() interface{} {
    if sll.IsEmpty(){
        return nil
    }
    first := sll.First()
    sll.head = sll.head.Next()
    sll.numItems--

    if sll.numItems == 0 {
        sll.tail = nil
    }

    return first
}

// Returns a string representation of the list
func (sll *SinglyLinkedList) String() string  {
    return fmt.Sprintf("<%v>", sll.numItems)
}
