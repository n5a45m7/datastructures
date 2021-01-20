package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceHeap(t *testing.T) {
	h := NewSliceHeap(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})

	// new heap is empty
	assert.True(t, h.IsEmpty())
	// empty heap panics on top
	assert.Panics(t, func() { h.Top() })
	// empty heap panics on pop
	assert.Panics(t, func() { h.Pop() })

	numbers := []int{4, 8, 1, 2, 5, 7, 6, 9, 3, 0}
	for _, number := range numbers {
		h.Push(number)
	}

	assert.Equal(t, 9, h.Top().(int))
	h.Pop()
	assert.Equal(t, 8, h.Top().(int))
	h.Pop()
	assert.Equal(t, 7, h.Top().(int))
	h.Pop()
	assert.Equal(t, 6, h.Top().(int))
	h.Pop()
	assert.Equal(t, 5, h.Top().(int))
	h.Pop()
	assert.Equal(t, 4, h.Top().(int))
	h.Pop()
	assert.Equal(t, 3, h.Top().(int))
	h.Pop()
	assert.Equal(t, 2, h.Top().(int))
	h.Pop()
	assert.Equal(t, 1, h.Top().(int))
	h.Pop()
	assert.Equal(t, 0, h.Top().(int))
	h.Pop()

	assert.True(t, h.IsEmpty())
}

func TestSliceHeapStructItem(t *testing.T) {
	type TeamResult struct {
		Name   string
		Points int
	}

	h := NewSliceHeap(func(a, b interface{}) bool {
		return a.(*TeamResult).Points > b.(*TeamResult).Points
	})

	// new heap is empty
	assert.True(t, h.IsEmpty())
	// empty heap panics on top
	assert.Panics(t, func() { h.Top() })
	// empty heap panics on pop
	assert.Panics(t, func() { h.Pop() })

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
		h.Push(team)
	}

	assert.Equal(t, "Blue", h.Top().(*TeamResult).Name)
	h.Pop()
	assert.Equal(t, "Red", h.Top().(*TeamResult).Name)
	h.Pop()
	assert.Equal(t, "Green", h.Top().(*TeamResult).Name)
	h.Pop()
	assert.Equal(t, "Yellow", h.Top().(*TeamResult).Name)
	h.Pop()

	assert.True(t, h.IsEmpty())
}
