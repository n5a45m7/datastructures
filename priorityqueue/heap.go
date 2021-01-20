package priorityqueue

import "n5a45m7/datastructures/heap"

// pqHeapBased is a heap based implementation of priority queue
type pqHeapBased struct {
	storage heap.Heap
}

// NewPQHeapBased create new instance of heap based priority queue
func NewPQHeapBased(hp HigherPriorityFunc) PriorityQueue {
	return &pqHeapBased{
		storage: heap.NewHeap(heap.HigherPriorityFunc(hp)),
	}
}

func (q *pqHeapBased) IsEmpty() bool {
	return q.storage.IsEmpty()
}

func (q *pqHeapBased) Top() interface{} {
	return q.storage.Top()
}
func (q *pqHeapBased) Pop() {
	q.storage.Pop()
}
func (q *pqHeapBased) Push(item interface{}) {
	q.storage.Push(item)
}
