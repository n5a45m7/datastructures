package stack

// NewLinkedListStack --
func NewLinkedListStack() Stack {
	return &linkedListStack{}
}

type node struct {
	value interface{}
	next  *node
}

type linkedListStack struct {
	top *node
}

func (s *linkedListStack) IsEmpty() bool {
	return s.top == nil
}

func (s *linkedListStack) Top() interface{} {
	return s.top.value
}

func (s *linkedListStack) Pop() {
	top := s.top
	s.top = s.top.next
	top.next = nil
}

func (s *linkedListStack) Push(item interface{}) {
	top := &node{
		value: item,
		next:  s.top,
	}
	s.top = top
}
