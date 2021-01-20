package heap

// HigherPriorityFunc returns true if and only if "a" HAS HIGHER PRIORITY THAN "b"
type HigherPriorityFunc func(a, b interface{}) bool

// Heap data structure
type Heap interface {
	IsEmpty() bool
	// schould panic on empty heap provide this check in your code
	Top() interface{}
	// schould panic on empty heap provide this check in your code
	Pop()
	Push(item interface{})
}

// NewHeap is a Heap creator function
func NewHeap(hp HigherPriorityFunc) Heap {
	return NewSliceHeap(hp)
}
