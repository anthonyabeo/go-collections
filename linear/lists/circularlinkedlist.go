package list

type CircularLinkedList struct {
	tail *node
	numItems int
}

func (cll CircularLinkedList) Size()  int {
	return cll.numItems
}

func (cll CircularLinkedList) IsEmpty() bool {
	return cll.numItems == 0
}

// returns the first element from the list
func (cll CircularLinkedList) First() interface{}  {
	if cll.IsEmpty(){
		return nil
	}
	return cll.tail.next.item
}

// returns the first element from the list
func (cll CircularLinkedList) Last() interface{} {
	if cll.IsEmpty(){
		return nil
	}
	return cll.tail.item
}


func (cll *CircularLinkedList) AddFirst(e interface{})  {
	if cll.IsEmpty() {
		cll.tail = &node{e, cll.tail}
		cll.tail.next = cll.tail
	} else {
		newNode := &node{e, cll.tail.next}
		cll.tail.next = newNode
	}
	cll.numItems++
}

func (cll *CircularLinkedList) AddLast(e interface{})  {
	cll.AddFirst(e)
	cll.tail = cll.tail.next
}

func (cll *CircularLinkedList) Rotate() {
	if cll.tail != nil {
		cll.tail = cll.tail.next
	}
}

func (cll *CircularLinkedList) RemoveFirst() interface{} {
	if cll.IsEmpty(){
		return nil
	}
	head := cll.tail.next
	if head == cll.tail {
		cll.tail = nil
	} else {
		cll.tail.next = head.next
	}
	cll.numItems--

	return head.item
}