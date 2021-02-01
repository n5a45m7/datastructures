package queue

type sliceQueue struct {
	storage []interface{}
}

// NewSliceQueue --
// initialCapacity - allocate storage slice with predefined capacity
func NewSliceQueue(initialCapacity uint) Queue {
	return &sliceQueue{
		storage: make([]interface{}, 0, initialCapacity),
	}
}

func (q *sliceQueue) IsEmpty() bool {
	return len(q.storage) <= 0
}

func (q *sliceQueue) Top() interface{} {
	return q.storage[0]
}

func (q *sliceQueue) Pop() {
	q.storage = q.storage[1:]
}

func (q *sliceQueue) Push(item interface{}) {
	q.storage = append(q.storage, item)
}
