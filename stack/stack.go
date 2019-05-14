package stack

import (
	"github.com/pkg/errors"

	"github.com/bitsgofer/containers"
)

// Stackable provides LIFO APIs.
type Stackable interface {
	Push(item containers.Value)
	Top() (containers.Value, error)
	Pop() (containers.Value, error)
	Size() int
	Clear()
}

// stack is a concrete implementation of Stackable on a slice.
type stack struct {
	items []containers.Value
}

func newZeroValueStack() *stack {
	return &stack{}
}

// New returns Stackable (with LIFO APIs).
func New() *stack {
	return newZeroValueStack()
}

// Push adds a new containers.Value on the stack's top.
func (s *stack) Push(item containers.Value) {
	s.items = append(s.items, item)
}

// Size returns the current number of elements in the stack.
func (s *stack) Size() int {
	return len(s.items)
}

// Clear empties the whole stack.
func (s *stack) Clear() {
	s.items = s.items[:0] // keep the unerlying array in heap
}

const stackIsEmpty = "stack is empty"

// Top returns the containers.Value on top of the stack.
func (s *stack) Top() (containers.Value, error) {
	n := len(s.items)
	if n == 0 {
		return nil, errors.New(stackIsEmpty)
	}

	return s.items[n-1], nil
}

// Pop removes the containers.Value on top of the stack and returns it.
func (s *stack) Pop() (containers.Value, error) {
	top, err := s.Top()
	if err != nil {
		return nil, err
	}

	s.items = s.items[:len(s.items)-1]
	return top, nil
}
