package list

import "fmt"

type dllNode struct {
	item interface{}
	next *dllNode
	prev *dllNode
}

// returns a pointer to the item in the node
func (n dllNode) Item() interface{} {
	return n.item
}

// returns a pointer to the next node in the list
func (n dllNode) Next() *dllNode {
	return n.next
}

// returns a pointer to the prev node in the list
func (n dllNode) Prev() *dllNode {
	return n.prev
}


type DoublyLinkedList struct {
	numItems int
	header *dllNode
	trailer *dllNode
}

func NewDll() *DoublyLinkedList {
	header := &dllNode{"header", nil, nil}
	trailer := &dllNode{"trailer", nil, header}

	dll := &DoublyLinkedList {
		0,
		header,
		trailer,
	}
	dll.header.next = trailer

	return dll
}

// Returns the number of items in the list
func (dll *DoublyLinkedList) Size() int {
	return dll.numItems
}

// returns a boolean indicating whether or not
// the list is empty or not
func (dll *DoublyLinkedList) IsEmpty() bool  {
	return dll.numItems == 0
}

// returns the first element from the list
func (dll *DoublyLinkedList) First() interface{}  {
	if dll.IsEmpty() {
		return nil
	}
	return dll.header.next.item
}

// returns the first element from the list
func (dll *DoublyLinkedList) Last() interface{} {
	if dll.IsEmpty() {
		return nil
	}
	return dll.trailer.prev.item
}

// Add a new item to the front of the list
func (dll *DoublyLinkedList) AddFirst(e interface{})  {
	head := dll.header
	headersNext := dll.header.next

	newNode := dllNode{e, headersNext, head}
	head.next = &newNode
	headersNext.prev = &newNode
	dll.numItems++
}

// Add a new item to the end of the list
func (dll *DoublyLinkedList) AddLast(e interface{})  {
	tail := dll.trailer
	trailerPrev := dll.trailer.prev

	newNode := dllNode{e, tail, trailerPrev}
	tail.prev = &newNode
	trailerPrev.next = &newNode
	dll.numItems++
}

// removes and returns the first item from the list
func (dll *DoublyLinkedList) RemoveFirst() interface{} {
	if dll.IsEmpty(){
		return nil
	}
	item := dll.header.next.item

	dll.header.next.next.prev = dll.header
	dll.header.next = dll.header.next.next
	dll.numItems--

	return item
}

// removes and returns the first item from the list
func (dll *DoublyLinkedList) RemoveLast() interface{} {
	if dll.IsEmpty(){
		return nil
	}

	item := dll.trailer.prev.item

	dll.trailer.prev.prev.next = dll.trailer
	dll.trailer.prev = dll.trailer.prev.prev
	dll.numItems--

	return item
}

func (dll *DoublyLinkedList) PrintForward()  {
	for curNode := dll.header.next; curNode.item != "trailer"; curNode = curNode.next {
		fmt.Println(curNode.item)
	}
}

func (dll *DoublyLinkedList) PrintReverse()  {
	for curNode := dll.trailer.prev; curNode.item != "header"; curNode = curNode.prev {
		fmt.Println(curNode.item)
	}
}