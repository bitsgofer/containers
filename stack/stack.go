package stack

import "github.com/pkg/errors"

// Element wraps a value in the stack.
type Element struct {
	Value interface{}
}

// Stackable provides LIFO APIs.
type Stackable interface {
	Push(item Element)
	Top() (Element, error)
	Pop() (Element, error)
	Size() int
	Clear()
}

// stack is a concrete implementation of Stackable on a slice.
type stack struct {
	items []Element
}

func newZeroValueStack() *stack {
	return &stack{}
}

// New returns Stackable (with LIFO APIs).
func New() *stack {
	return newZeroValueStack()
}

// Push adds a new Element on the stack's top.
func (s *stack) Push(item Element) {
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

// Top returns the Element on top of the stack.
func (s *stack) Top() (Element, error) {
	n := len(s.items)
	if n == 0 {
		return Element{}, errors.New(stackIsEmpty)
	}

	return s.items[n-1], nil
}

// Pop removes the Element on top of the stack and returns it.
func (s *stack) Pop() (Element, error) {
	top, err := s.Top()
	if err != nil {
		return Element{}, err
	}

	s.items = s.items[:len(s.items)-1]
	return top, nil
}
