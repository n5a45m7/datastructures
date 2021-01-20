package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapPriorityQueue(t *testing.T) {
	q := NewPQHeapBased(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})

	// new heap is empty
	assert.True(t, q.IsEmpty())
	// empty heap panics on top
	assert.Panics(t, func() { q.Top() })
	// empty heap panics on pop
	assert.Panics(t, func() { q.Pop() })

	numbers := []int{4, 8, 1, 2, 5, 7, 6, 9, 3, 0}
	for _, number := range numbers {
		q.Push(number)
	}

	assert.Equal(t, 9, q.Top().(int))
	q.Pop()
	assert.Equal(t, 8, q.Top().(int))
	q.Pop()
	assert.Equal(t, 7, q.Top().(int))
	q.Pop()
	assert.Equal(t, 6, q.Top().(int))
	q.Pop()
	assert.Equal(t, 5, q.Top().(int))
	q.Pop()
	assert.Equal(t, 4, q.Top().(int))
	q.Pop()
	assert.Equal(t, 3, q.Top().(int))
	q.Pop()
	assert.Equal(t, 2, q.Top().(int))
	q.Pop()
	assert.Equal(t, 1, q.Top().(int))
	q.Pop()
	assert.Equal(t, 0, q.Top().(int))
	q.Pop()

	assert.True(t, q.IsEmpty())
}

func TestHeapPriorityQueueStructItem(t *testing.T) {
	type TeamResult struct {
		Name   string
		Points int
	}

	q := NewPQHeapBased(func(a, b interface{}) bool {
		return a.(*TeamResult).Points > b.(*TeamResult).Points
	})

	// new heap is empty
	assert.True(t, q.IsEmpty())
	// empty heap panics on top
	assert.Panics(t, func() { q.Top() })
	// empty heap panics on pop
	assert.Panics(t, func() { q.Pop() })

	teams := []*TeamResult{
		&TeamResult{
			Name:   "Red",
			Points: 100,
		},
		&TeamResult{
			Name:   "Blue",
			Points: 109,
		},
		&TeamResult{
			Name:   "Yellow",
			Points: 84,
		},
		&TeamResult{
			Name:   "Green",
			Points: 89,
		},
	}
	for _, team := range teams {
		q.Push(team)
	}

	assert.Equal(t, "Blue", q.Top().(*TeamResult).Name)
	q.Pop()
	assert.Equal(t, "Red", q.Top().(*TeamResult).Name)
	q.Pop()
	assert.Equal(t, "Green", q.Top().(*TeamResult).Name)
	q.Pop()
	assert.Equal(t, "Yellow", q.Top().(*TeamResult).Name)
	q.Pop()

	assert.True(t, q.IsEmpty())
}
