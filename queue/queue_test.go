package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testQueue(q Queue) func(t *testing.T) {
	return func(t *testing.T) {
		assert.True(t, q.IsEmpty())
		// empty queue panics on top
		assert.Panics(t, func() { q.Top() })
		// empty queue panics on pop
		assert.Panics(t, func() { q.Top() })

		numbers := []int{5, 1, 0, 2, 3, 7, 9, 4, 8, 6}
		for _, number := range numbers {
			q.Push(number)
		}

		assert.False(t, q.IsEmpty())
		assert.Equal(t, numbers[0], q.Top().(int))
		assert.Equal(t, numbers[0], q.Top().(int))
		for _, number := range numbers {
			assert.False(t, q.IsEmpty())
			assert.Equal(t, number, q.Top().(int))
			q.Pop()
		}
		assert.True(t, q.IsEmpty())
	}
}
