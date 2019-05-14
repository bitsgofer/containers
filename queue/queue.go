package queue

import (
	"github.com/pkg/errors"

	"github.com/bitsgofer/containers"
)

// Queueable provides LIFO APIs.
type Queueable interface {
	Enqueue(item containers.Value)
	Dequeue() (containers.Value, error)
	Front() (containers.Value, error)
	Back() (containers.Value, error)
	Size() int
	Clear()
}

// queue is a concrete implementation of Queueable on a slice.
type queue struct {
	items []containers.Value
}

func newZeroValueQueue() *queue {
	return &queue{}
}

// New returns Queueable (with FIFO APIs).
func New() *queue {
	return newZeroValueQueue()
}

// Enqueue adds a new containers.Value on the queue's top.
func (s *queue) Enqueue(item containers.Value) {
	s.items = append(s.items, item)
}

// Size returns the current number of elements in the queue.
func (s *queue) Size() int {
	return len(s.items)
}

// Clear empties the whole queue.
func (s *queue) Clear() {
	s.items = s.items[:0] // keep the unerlying array in heap
}

const queueIsEmpty = "queue is empty"

// Front returns the containers.Value on top of the queue.
func (s *queue) Front() (containers.Value, error) {
	n := len(s.items)
	if n == 0 {
		return nil, errors.New(queueIsEmpty)
	}

	return s.items[0], nil
}

// Back returns the containers.Value on top of the queue.
func (s *queue) Back() (containers.Value, error) {
	n := len(s.items)
	if n == 0 {
		return nil, errors.New(queueIsEmpty)
	}

	return s.items[n-1], nil
}

// Dequeue removes the containers.Value on top of the queue and returns it.
func (s *queue) Dequeue() (containers.Value, error) {
	top, err := s.Front()
	if err != nil {
		return nil, err
	}

	s.items = s.items[1:]
	return top, nil
}
