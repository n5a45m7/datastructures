package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	s := NewSliceStack(0)
	testStack(s)(t)

	s = NewSliceStack(0)
	testStackStructItem(s)(t)
}
