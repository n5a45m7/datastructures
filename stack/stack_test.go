package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testStack(s Stack) func(t *testing.T) {
	return func(t *testing.T) {
		assert.True(t, s.IsEmpty())
		// empty stack panics on top
		assert.Panics(t, func() { s.Top() })
		// empty stack panics on pop
		assert.Panics(t, func() { s.Top() })

		numbers := []int{5, 1, 0, 2, 3, 7, 9, 4, 8, 6}
		for _, number := range numbers {
			s.Push(number)
		}

		assert.False(t, s.IsEmpty())
		assert.Equal(t, numbers[len(numbers)-1], s.Top().(int))
		assert.Equal(t, numbers[len(numbers)-1], s.Top().(int))
		for i := len(numbers) - 1; i >= 0; i-- {
			assert.False(t, s.IsEmpty())
			assert.Equal(t, numbers[i], s.Top().(int))
			s.Pop()
		}
		assert.True(t, s.IsEmpty())
	}
}

func testStackStructItem(s Stack) func(t *testing.T) {
	return func(t *testing.T) {
		type stackItem struct {
			value int
		}
		assert.True(t, s.IsEmpty())
		// empty stack panics on top
		assert.Panics(t, func() { s.Top() })
		// empty stack panics on pop
		assert.Panics(t, func() { s.Top() })

		items := []stackItem{
			stackItem{4},
			stackItem{2},
			stackItem{3},
			stackItem{0},
			stackItem{1},
		}
		for _, item := range items {
			s.Push(item)
		}

		assert.False(t, s.IsEmpty())
		assert.Equal(t, items[len(items)-1], s.Top().(stackItem))
		assert.Equal(t, items[len(items)-1], s.Top().(stackItem))
		for i := len(items) - 1; i >= 0; i-- {
			assert.False(t, s.IsEmpty())
			assert.Equal(t, items[i], s.Top().(stackItem))
			s.Pop()
		}
		assert.True(t, s.IsEmpty())
	}
}
