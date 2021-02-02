package stack

// NewSliceStack --
// initialCapacity - allocate storage slice with predefined capacity
func NewSliceStack(initialCapacity int) Stack {
	return &sliceStack{
		storage: make([]interface{}, 0, initialCapacity),
	}
}

type sliceStack struct {
	storage []interface{}
}

func (s *sliceStack) IsEmpty() bool {
	return len(s.storage) <= 0
}

func (s *sliceStack) Top() interface{} {
	return s.storage[len(s.storage)-1]
}

func (s *sliceStack) Pop() {
	s.storage = s.storage[:len(s.storage)-1]
}

func (s *sliceStack) Push(item interface{}) {
	s.storage = append(s.storage, item)
}
