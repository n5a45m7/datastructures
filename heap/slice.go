package heap

// slice is a slice based implementation of Heap interface
type slice struct {
	storage []interface{}
	hp      HigherPriorityFunc
}

// NewSliceHeap creates slice entity
// hp returns true if and only if "a" HAS HIGHER PRIORITY THAN "b"
func NewSliceHeap(hp HigherPriorityFunc) Heap {
	return &slice{
		storage: make([]interface{}, 0),
		hp:      hp,
	}
}

// IsEmpty --
func (h *slice) IsEmpty() bool {
	return len(h.storage) <= 0
}

// Top --
func (h *slice) Top() interface{} {
	return h.storage[0]
}

// Pop --
func (h *slice) Pop() {
	// last element index
	li := len(h.storage) - 1
	// swap
	h.storage[0], h.storage[li] = h.storage[li], h.storage[0]
	// truncate last element of slice
	h.storage = h.storage[:li]
	h.heapify(0)
}

// Push --
func (h *slice) Push(item interface{}) {
	h.storage = append(h.storage, item)
	h.upwards(len(h.storage) - 1)
}

// heapify from index i means
// move element down in the tree based on priority value
func (h *slice) heapify(i int) {
	//ci is a child indexes, slice of child nodes index
	// 2*i+1 is index of left child node for node i
	// 2*i+2 is index of right child node for node i
	ci := []int{2*i + 1, 2*i + 2}
	// index of node with top priority, init with current node i
	tpi := i

	// ci is a child index
	for _, j := range ci {
		if j < len(h.storage) && h.hp(h.storage[j], h.storage[tpi]) {
			tpi = j
		}
	}

	// if there is a child node which has more valuable priority than i, swap and heapify further
	if tpi != i {
		h.storage[i], h.storage[tpi] = h.storage[tpi], h.storage[i]
		h.heapify(tpi)
	}
}

// upwards from index i means
// move element up in tree based on priority value
func (h *slice) upwards(i int) {
	if i <= 0 {
		return
	}
	// pi is a parent index
	pi := (i - 1) / 2
	// if priority of node is is more vluable than its parent, swap and check further for parent node
	if h.hp(h.storage[i], h.storage[pi]) {
		h.storage[i], h.storage[pi] = h.storage[pi], h.storage[i]
		h.upwards(pi)
	}
}
