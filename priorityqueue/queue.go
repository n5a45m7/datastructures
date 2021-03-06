package priorityqueue

import "github.com/n5a45m7/datastructures/heap"

// HigherPriorityFunc returns true if and only if "a" HAS HIGHER PRIORITY THAN "b"
type HigherPriorityFunc func(a, b interface{}) bool

// PriorityQueue data structure
type PriorityQueue interface {
	IsEmpty() bool
	// schould panic on empty queue provide this check in your code
	Top() interface{}
	// schould panic on empty queue provide this check in your code
	Pop()
	Push(item interface{})
}

// NewPriorityQueue is a Queue creator function
func NewPriorityQueue(hp HigherPriorityFunc) PriorityQueue {
	// heap based priority queue
	return heap.NewHeap(heap.HigherPriorityFunc(hp))
}
