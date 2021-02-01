package queue

// Queue --
type Queue interface {
	IsEmpty() bool
	// schould panic on empty queue provide this check in your code
	Top() interface{}
	// schould panic on empty queue provide this check in your code
	Pop()
	Push(item interface{})
}

// NewQueue --
func NewQueue() Queue {
	return NewSliceQueue(32)
}
