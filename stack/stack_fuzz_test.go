// +build fuzz

package stack

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bitsgofer/containers"
)

// TestFuzzOps performs N random operations to see if the stack panic.
func TestFuzzOps(t *testing.T) {
	randSeed := time.Now().Unix()
	rng := rand.New(rand.NewSource(randSeed))
	t.Logf("running with random seed= %v", randSeed)

	s := newZeroValueStack()
	push := func() {
		s.Push(containers.Value(rng.Int()))
	}
	pop := func() { s.Pop() }
	top := func() { s.Top() }
	size := func() { s.Size() }
	clear := func() { s.Clear() }

	steps := rng.Intn(10000) + 2000
	for i := 0; i < steps; i++ {
		switch rng.Intn(5) {
		case 0:
			push()
		case 1:
			pop()
		case 2:
			top()
		case 3:
			size()
		default:
			clear()
		}
	}
}
