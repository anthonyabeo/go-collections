package linear

import "go-collections/linear/lists/linked"

type Queue struct {
	queue linked.SinglyLinkedList
}

func (q *Queue) Enqueue(e interface{}) {
	q.queue.AddLast(e)
}

func (q *Queue) Dequeue() interface{} {
	return q.queue.RemoveFirst()
}

func (q *Queue) First() interface{} {
	return q.queue.First()
}

func (q *Queue) Size() int {
	return q.queue.Size()
}

func (q *Queue) IsEmpty() bool {
	return q.queue.IsEmpty()
}
