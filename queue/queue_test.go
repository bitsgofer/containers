package queue

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/bitsgofer/containers"
)

var queueCmpOption = cmp.AllowUnexported(queue{})

func TestEnqueue(t *testing.T) {
	var testCases = map[string]struct {
		s         *queue
		item      containers.Value
		nextQueue *queue
	}{
		"zeroValue": {
			s:         newZeroValueQueue(),
			item:      containers.Value(1),
			nextQueue: newQueueWithValues(1),
		},
		"empty": {
			s:         newQueueWithValues(),
			item:      containers.Value(1),
			nextQueue: newQueueWithValues(1),
		},
		"filled": {
			s:         newQueueWithValues(1, 2, 3),
			item:      containers.Value(4),
			nextQueue: newQueueWithValues(1, 2, 3, 4),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.s.Enqueue(tc.item)

			if want, got := tc.nextQueue, tc.s; !cmp.Equal(want, got, queueCmpOption) {
				t.Fatalf("want next queue= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got, queueCmpOption))
			}
		})
	}
}

func TestClear(t *testing.T) {
	var testCases = map[string]struct {
		s         *queue
		nextQueue *queue
	}{
		"zeroValue": {
			s:         newZeroValueQueue(),
			nextQueue: newZeroValueQueue(),
		},
		"empty": {
			s:         newQueueWithValues(),
			nextQueue: newQueueWithValues(),
		},
		"filled": {
			s:         newQueueWithValues(1, 2, 3),
			nextQueue: newQueueWithValues(),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.s.Clear()

			if want, got := tc.nextQueue, tc.s; !cmp.Equal(want, got, queueCmpOption) {
				t.Fatalf("want next queue= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got, queueCmpOption))
			}
		})
	}
}

func TestSize(t *testing.T) {
	var testCases = map[string]struct {
		s    *queue
		size int
	}{
		"zeroValue": {
			s:    newZeroValueQueue(),
			size: 0,
		},
		"empty": {
			s:    newQueueWithValues(),
			size: 0,
		},
		"one": {
			s:    newQueueWithValues(3),
			size: 1,
		},
		"many": {
			s:    newQueueWithValues(1, 2, 3),
			size: 3,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if want, got := tc.size, tc.s.Size(); want != got {
				t.Fatalf("want= %v, got= %v", want, got)
			}
		})
	}
}

func TestFront(t *testing.T) {
	var testCases = map[string]struct {
		s     *queue
		isErr bool
		val   containers.Value
	}{
		"zeroValue": {
			s:     newZeroValueQueue(),
			isErr: true,
		},
		"empty": {
			s:     newQueueWithValues(),
			isErr: true,
		},
		"filled": {
			s:   newQueueWithValues(1, 2, 3),
			val: containers.Value(1),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			val, err := tc.s.Front()

			if tc.isErr && err == nil {
				t.Fatalf("want error, got none")
			}
			if !tc.isErr && err != nil {
				t.Fatalf("want no error, got %q", err)
			}
			if want, got := tc.val, val; !cmp.Equal(want, got) {
				t.Fatalf("want %v, got %v, diff= %v", want, got, cmp.Diff(want, got))
			}
		})
	}
}

func TestBack(t *testing.T) {
	var testCases = map[string]struct {
		s     *queue
		isErr bool
		val   containers.Value
	}{
		"zeroValue": {
			s:     newZeroValueQueue(),
			isErr: true,
		},
		"empty": {
			s:     newQueueWithValues(),
			isErr: true,
		},
		"filled": {
			s:   newQueueWithValues(1, 2, 3),
			val: containers.Value(3),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			val, err := tc.s.Back()

			if tc.isErr && err == nil {
				t.Fatalf("want error, got none")
			}
			if !tc.isErr && err != nil {
				t.Fatalf("want no error, got %q", err)
			}
			if want, got := tc.val, val; !cmp.Equal(want, got) {
				t.Fatalf("want %v, got %v, diff= %v", want, got, cmp.Diff(want, got))
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	var testCases = map[string]struct {
		s                 *queue
		isErr             bool
		val               containers.Value
		emptyAfterDequeue bool
		nextQueue         *queue
	}{
		"zeroValue": {
			s:         newZeroValueQueue(),
			isErr:     true,
			nextQueue: newZeroValueQueue(),
		},
		"empty": {
			s:         newQueueWithValues(),
			isErr:     true,
			nextQueue: newQueueWithValues(),
		},
		"oneElement": {
			s:         newQueueWithValues(1),
			val:       containers.Value(1),
			nextQueue: newQueueWithValues(),
		},
		"manyElements": {
			s:         newQueueWithValues(1, 2, 3),
			val:       containers.Value(1),
			nextQueue: newQueueWithValues(2, 3),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			val, err := tc.s.Dequeue()

			switch {
			case tc.isErr && err == nil:
				t.Fatalf("want error, got none")
			case !tc.isErr && err != nil:
				t.Fatalf("want no error, got %q", err)
			default:
				if want, got := tc.val, val; !cmp.Equal(want, got) {
					t.Fatalf("want= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got))
				}
			}

			if want, got := tc.nextQueue, tc.s; !cmp.Equal(want, got, queueCmpOption) {
				t.Fatalf("want next queue= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got, queueCmpOption))
			}
		})
	}
}

func newQueueWithValues(vals ...interface{}) *queue {
	s := &queue{
		items: make([]containers.Value, 0, len(vals)),
	}
	for _, v := range vals {
		s.Enqueue(containers.Value(v))
	}

	return s
}

var benchmarkTypes = map[string]struct {
	newValue func() containers.Value
}{
	"int": {
		newValue: func() containers.Value {
			return containers.Value(2)
		},
	},
	"string": {
		newValue: func() containers.Value {
			return containers.Value("this is a string")
		},
	},
	"largeStruct": {
		newValue: func() containers.Value {
			return containers.Value(http.Client{})
		},
	},
}

func BenchmarkEnqueue(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueQueue()
			for i := 0; i < b.N; i++ {
				s.Enqueue(bm.newValue())
			}
		})
	}
}

func BenchmarkFront(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueQueue()
			for i := 0; i < 10; i++ {
				s.Enqueue(bm.newValue())
			}

			for i := 0; i < b.N; i++ {
				s.Front()
			}
		})
	}
}

func BenchmarkBack(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueQueue()
			for i := 0; i < 10; i++ {
				s.Enqueue(bm.newValue())
			}

			for i := 0; i < b.N; i++ {
				s.Back()
			}
		})
	}
}

func BenchmarkDequeue(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueQueue()
			for i := 0; i < 10; i++ {
				s.Enqueue(bm.newValue())
			}

			for i := 0; i < b.N; i++ {
				s.Dequeue()
			}
		})
	}
}
