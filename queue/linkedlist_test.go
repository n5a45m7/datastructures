package queue

import (
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	testQueue(q)(t)
}
