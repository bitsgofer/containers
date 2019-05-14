package btree

import (
	"math"
	"testing"

	"github.com/bitsgofer/containers"
)

func TestBST(t *testing.T) {
	var testCases = map[string]struct {
		vals  []containers.Value
		isBST bool
	}{
		"rootOnly": {
			vals:  []containers.Value{1},
			isBST: true,
		},
		"twoLvlFullBST": {
			vals:  []containers.Value{2, 1, 3},
			isBST: true,
		},
		"rightTreeElementLessThanRoot": {
			vals:  []containers.Value{5, 1, 4, nil, nil, 3, 6},
			isBST: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			root, err := NewFromLeetCodeOrder(tc.vals...)
			if err != nil {
				t.Fatalf("must give a binary tree")
			}

			less := func(a, b containers.Value) bool {
				return a.(int) < b.(int)
			}
			if want, got := tc.isBST, IsBST(root, less, math.MinInt64, math.MaxInt64); want != got {
				t.Fatalf("want= %v, got= %v", want, got)
			}
		})
	}
}
