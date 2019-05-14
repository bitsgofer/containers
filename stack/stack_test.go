package stack

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/bitsgofer/containers"
)

var stackCmpOption = cmp.AllowUnexported(stack{})

func TestPush(t *testing.T) {
	var testCases = map[string]struct {
		s         *stack
		item      containers.Value
		nextStack *stack
	}{
		"zeroValue": {
			s:         newZeroValueStack(),
			item:      containers.Value(1),
			nextStack: newStackWithValues(1),
		},
		"empty": {
			s:         newStackWithValues(),
			item:      containers.Value(1),
			nextStack: newStackWithValues(1),
		},
		"filled": {
			s:         newStackWithValues(1, 2, 3),
			item:      containers.Value(4),
			nextStack: newStackWithValues(1, 2, 3, 4),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.s.Push(tc.item)

			if want, got := tc.nextStack, tc.s; !cmp.Equal(want, got, stackCmpOption) {
				t.Fatalf("want next stack= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got, stackCmpOption))
			}
		})
	}
}

func TestClear(t *testing.T) {
	var testCases = map[string]struct {
		s         *stack
		nextStack *stack
	}{
		"zeroValue": {
			s:         newZeroValueStack(),
			nextStack: newZeroValueStack(),
		},
		"empty": {
			s:         newStackWithValues(),
			nextStack: newStackWithValues(),
		},
		"filled": {
			s:         newStackWithValues(1, 2, 3),
			nextStack: newStackWithValues(),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.s.Clear()

			if want, got := tc.nextStack, tc.s; !cmp.Equal(want, got, stackCmpOption) {
				t.Fatalf("want next stack= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got, stackCmpOption))
			}
		})
	}
}

func TestSize(t *testing.T) {
	var testCases = map[string]struct {
		s    *stack
		size int
	}{
		"zeroValue": {
			s:    newZeroValueStack(),
			size: 0,
		},
		"empty": {
			s:    newStackWithValues(),
			size: 0,
		},
		"one": {
			s:    newStackWithValues(3),
			size: 1,
		},
		"many": {
			s:    newStackWithValues(1, 2, 3),
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

func TestTop(t *testing.T) {
	var testCases = map[string]struct {
		s     *stack
		isErr bool
		val   containers.Value
	}{
		"zeroValue": {
			s:     newZeroValueStack(),
			isErr: true,
		},
		"empty": {
			s:     newStackWithValues(),
			isErr: true,
		},
		"filled": {
			s:   newStackWithValues(1, 2, 3),
			val: containers.Value(3),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			val, err := tc.s.Top()

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

func TestPop(t *testing.T) {
	var testCases = map[string]struct {
		s             *stack
		isErr         bool
		val           containers.Value
		emptyAfterPop bool
		nextStack     *stack
	}{
		"zeroValue": {
			s:         newZeroValueStack(),
			isErr:     true,
			nextStack: newZeroValueStack(),
		},
		"empty": {
			s:         newStackWithValues(),
			isErr:     true,
			nextStack: newStackWithValues(),
		},
		"oneElement": {
			s:         newStackWithValues(1),
			val:       containers.Value(1),
			nextStack: newStackWithValues(),
		},
		"manyElements": {
			s:         newStackWithValues(1, 2, 3),
			val:       containers.Value(3),
			nextStack: newStackWithValues(1, 2),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			val, err := tc.s.Pop()

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

			if want, got := tc.nextStack, tc.s; !cmp.Equal(want, got, stackCmpOption) {
				t.Fatalf("want next stack= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got, stackCmpOption))
			}
		})
	}
}

func newStackWithValues(vals ...interface{}) *stack {
	s := &stack{
		items: make([]containers.Value, 0, len(vals)),
	}
	for _, v := range vals {
		s.Push(containers.Value(v))
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

func BenchmarkPush(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueStack()
			for i := 0; i < b.N; i++ {
				s.Push(bm.newValue())
			}
		})
	}
}

func BenchmarkTop(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueStack()
			for i := 0; i < 10; i++ {
				s.Push(bm.newValue())
			}

			for i := 0; i < b.N; i++ {
				s.Top()
			}
		})
	}
}

func BenchmarkPop(b *testing.B) {
	for name, bm := range benchmarkTypes {
		b.Run(name, func(b *testing.B) {
			s := newZeroValueStack()
			for i := 0; i < 10; i++ {
				s.Push(bm.newValue())
			}

			for i := 0; i < b.N; i++ {
				s.Pop()
			}
		})
	}
}
