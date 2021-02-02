package stack

// Stack --
type Stack interface {
	IsEmpty() bool
	// schould panic on empty stack provide this check in your code
	Top() interface{}
	// schould panic on empty stack provide this check in your code
	Pop()
	Push(item interface{})
}

// NewStack --
func NewStack() Stack {
	return NewSliceStack(32)
}
