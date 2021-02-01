package queue

type node struct {
	next  *node
	value interface{}
}

type linkedListQueue struct {
	head *node
	tail *node
}

// NewLinkedListQueue --
func NewLinkedListQueue() Queue {
	return &linkedListQueue{}
}

func (q *linkedListQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *linkedListQueue) Top() interface{} {
	return q.head.value
}

func (q *linkedListQueue) Pop() {
	n := q.head
	q.head = q.head.next
	n.next = nil

	if q.head == nil {
		q.tail = nil
	}
}

func (q *linkedListQueue) Push(item interface{}) {
	n := &node{value: item}
	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}

	q.tail.next = n
	q.tail = n
}
