package stack

import (
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	s := NewLinkedListStack()
	testStack(s)(t)

	s = NewLinkedListStack()
	testStackStructItem(s)(t)
}
