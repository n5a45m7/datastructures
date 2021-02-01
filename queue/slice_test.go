package queue

import (
	"testing"
)

func TestSliceQueue(t *testing.T) {
	q := NewSliceQueue(0)
	testQueue(q)(t)
}
