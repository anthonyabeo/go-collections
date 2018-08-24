package linear

import "go-collections/linear/lists/linked"

type Deque struct {
	deque *linked.DoublyLinkedList
}

func NewDeque() *Deque {
	return &Deque{linked.NewDll()}
}

func (deque *Deque) AddFirst(e interface{}) {
	deque.deque.AddFirst(e)
}

func (deque *Deque) AddLast(e interface{}) {
	deque.deque.AddLast(e)
}

func (deque *Deque) RemoveFirst() interface{} {
	return deque.deque.RemoveFirst()
}

func (deque *Deque) RemoveLast() interface{} {
	return deque.deque.RemoveLast()
}

func (deque *Deque) Print() {
	deque.deque.PrintForward()
}

func (deque *Deque) First() interface{} {
	return deque.deque.First()
}

func (deque *Deque) Last() interface{} {
	return deque.deque.Last()
}

func (deque *Deque) Size() int {
	return deque.deque.Size()
}

func (deque *Deque) IsEmpty() bool {
	return deque.deque.IsEmpty()
}
