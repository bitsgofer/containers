// +build fuzz

package queue

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bitsgofer/containers"
)

// TestFuzzOps performs N random operations to see if the queue panic.
func TestFuzzOps(t *testing.T) {
	randSeed := time.Now().Unix()
	rng := rand.New(rand.NewSource(randSeed))
	t.Logf("running with random seed= %v", randSeed)

	s := newZeroValueQueue()
	enqueue := func() {
		s.Enqueue(containers.Value(rng.Int()))
	}
	dequeue := func() { s.Dequeue() }
	front := func() { s.Front() }
	back := func() { s.Back() }
	size := func() { s.Size() }
	clear := func() { s.Clear() }

	steps := rng.Intn(10000) + 2000
	for i := 0; i < steps; i++ {
		switch rng.Intn(6) {
		case 0:
			enqueue()
		case 1:
			dequeue()
		case 2:
			front()
		case 3:
			back()
		case 4:
			size()
		default:
			clear()
		}
	}
}
