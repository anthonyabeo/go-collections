package list

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
func (dll DoublyLinkedList) Size() int {
	return dll.numItems
}

// returns a boolean indicating whether or not
// the list is empty or not
func (dll DoublyLinkedList) IsEmpty() bool  {
	return dll.numItems == 0
}

// returns the first element from the list
func (dll DoublyLinkedList) First() interface{}  {
	if dll.IsEmpty() {
		return nil
	}
	return dll.header.next.item
}

// returns the first element from the list
func (dll DoublyLinkedList) Last() interface{} {
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

//// Add a new item to the end of the list
//func (dll DoublyLinkedList) AddLast(e interface{})  {
//
//}

//// removes and returns the first item from the list
//func (dll DoublyLinkedList) RemoveFirst() interface{} {
//
//}

// removes and returns the first item from the list
//func (dll DoublyLinkedList) RemoveLast() interface{} {
//	return interface()
//}