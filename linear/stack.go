package linear

import "go-collections/linear/lists"

type Stack struct {
	stack list.SinglyLinkedList
}

func (stack *Stack) Push(e interface{})  {
	stack.stack.AddFirst(e)
}

func (stack *Stack) Pop() interface{} {
	return stack.stack.RemoveFirst()
}

func (stack *Stack) Top() interface{} {
	return stack.stack.First()
}

func (stack *Stack) Size() int {
	return stack.stack.Size()
}

func (stack *Stack) IsEmpty() bool {
	return stack.stack.IsEmpty()
}
