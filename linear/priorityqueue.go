package linear

type Entry struct {
	key   interface{}
	value interface{}
}

func (entry Entry) Parent(j int) int   { return (j - 1) / 2 }
func (entry Entry) Left(j int) int     { return (2 * j) + 1 }
func (entry Entry) Right(j int) int    { return (2 * j) + 2 }
func (entry Entry) Key() interface{}   { return entry.key }
func (entry Entry) Value() interface{} { return entry.value }
func (entry *Entry) compare(other interface{}) int {
	var x int
	switch other.(type) {
	case int:
		if entry.key.(int) == other.(int) {
			x = 0
		} else if entry.key.(int) > other.(int) {
			x = 1
		} else {
			x = -1
		}
	case string:
		if entry.key.(string) == other.(string) {
			x = 0
		} else if entry.key.(string) > other.(string) {
			x = 1
		} else {
			x = -1
		}
	case float32:
		if entry.key.(float32) == other.(float32) {
			x = 0
		} else if entry.key.(float32) > other.(float32) {
			x = 1
		} else {
			x = -1
		}
	case uint:
		if entry.key.(uint) == other.(uint) {
			x = 0
		} else if entry.key.(uint) > other.(uint) {
			x = 1
		} else {
			x = -1
		}
	case rune:
		if entry.key.(rune) == other.(rune) {
			x = 0
		} else if entry.key.(rune) > other.(rune) {
			x = 1
		} else {
			x = -1
		}
	case byte:
		if entry.key.(byte) == other.(byte) {
			x = 0
		} else if entry.key.(byte) > other.(byte) {
			x = 1
		} else {
			x = -1
		}
	}

	return x
}

type PriorityQueue struct {
	heap []interface{}
	size int
}

func NewPQ() *PriorityQueue {
	return &PriorityQueue{
		make([]interface{}, 10),
		0,
	}
}

//
func (pq *PriorityQueue) Heap() []interface{} { return pq.heap }

//
func (pq *PriorityQueue) IsEmpty() bool { return pq.size == 0 }

//
func (pq *PriorityQueue) Size() int { return pq.size }

//
func (pq *PriorityQueue) Insert(key interface{}, value interface{}) {
	entry := &Entry{key, value}
	entryIdx := pq.size

	if pq.heap[entryIdx] != nil {
		return
	}
	pq.heap[entryIdx] = entry

	for entryIdx > 0 {
		parentIdx := entry.Parent(entryIdx)
		parent := pq.heap[parentIdx].(*Entry)

		if entry.compare(parent.key) > 0 {
			break
		}
		swap(pq, parentIdx, entry, entryIdx)
		entryIdx = parentIdx
	}
	pq.size++
}

//
func (pq *PriorityQueue) Min() *Entry {
	if pq.heap[0] == nil {
		return nil
	}
	return pq.heap[0].(*Entry)
}

//
func (pq *PriorityQueue) RemoveMin() *Entry {
	if pq.IsEmpty() {
		return nil
	}

	delItem := pq.heap[pq.size-1].(*Entry)

	delItemIdx := 0
	retVal := pq.heap[delItemIdx].(*Entry)
	pq.heap[delItemIdx] = delItem

	for (2*delItemIdx)+1 < pq.size {
		smallerChild, smallerChildIdx := getSmallerChild(pq, delItemIdx)
		if smallerChild.compare(delItem.key) > 0 {
			break
		}

		swap(pq, delItemIdx, smallerChild, smallerChildIdx)
		delItemIdx = smallerChildIdx
	}
	pq.size--
	pq.heap[pq.size] = nil

	return retVal
}

func swap(pq *PriorityQueue, first int, entry *Entry, entryIdx int) {
	temp := pq.heap[first]
	pq.heap[first] = entry
	pq.heap[entryIdx] = temp
}

func getSmallerChild(pq *PriorityQueue, idx int) (*Entry, int) {

	lastItemLeftChild := pq.heap[(2*idx)+1].(*Entry)
	lastItemRightChild := pq.heap[(2*idx)+2].(*Entry)
	var child *Entry
	var childIdx int

	if lastItemLeftChild.compare(lastItemRightChild.key) < 0 {
		childIdx = (2 * idx) + 1
		child = lastItemLeftChild
	} else {
		childIdx = (2 * idx) + 2
		child = lastItemRightChild
	}

	return child, childIdx
}
