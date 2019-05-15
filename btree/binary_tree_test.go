package btree

import (
	"testing"

	"github.com/bitsgofer/containers"
	"github.com/google/go-cmp/cmp"
)

func TestNewFromLeetCoderder(t *testing.T) {
	var testCases = map[string]struct {
		vals        []containers.Value
		isErr       bool
		valsInOrder []containers.Value
	}{
		"empty": {
			vals:  nil,
			isErr: true,
		},
		"oneElement": {
			vals:        []containers.Value{1},
			valsInOrder: []containers.Value{1},
		},
		"oneLvlFull": {
			vals:        []containers.Value{1, 2, 3},
			valsInOrder: []containers.Value{1, 2, 3},
		},
		"threeLvlWithNil": {
			vals:        []containers.Value{1, nil, 3, nil, nil, 6, 7, nil, nil, nil, nil, 11},
			valsInOrder: []containers.Value{1, 3, 6, 11, 7},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			root, err := NewFromLeetCodeOrder(tc.vals...)

			switch {
			case tc.isErr && err == nil:
				t.Fatalf("want error, got none")
			case !tc.isErr && err != nil:
				t.Fatalf("want no error, got %v", err)
			case tc.isErr && err != nil:
				return
			default:
				if want, got := tc.valsInOrder, extracValuesPreOrder(root); !cmp.Equal(want, got) {
					t.Fatalf("values inorder: want %v, got %v, diff= %v", want, got, cmp.Diff(want, got))
				}
			}
		})
	}
}

func TestExtractValuesPreOrder(t *testing.T) {
	var testCases = map[string]struct {
		vals         []containers.Value
		valsPreOrder []containers.Value
	}{
		// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder
		"geeksForGeeks": {
			vals:         []containers.Value{1, 2, 3, 4, 5},
			valsPreOrder: []containers.Value{1, 2, 4, 5, 3},
		},
		// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/992
		"leetCode": {
			vals:         []containers.Value{"F", "B", "G", "A", "D", nil, "I", nil, nil, "C", "E", nil, nil, "H"},
			valsPreOrder: []containers.Value{"F", "B", "A", "D", "C", "E", "G", "I", "H"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			root, err := NewFromLeetCodeOrder(tc.vals...)
			if err != nil {
				t.Fatalf("need a valid btree to test")
			}

			if want, got := tc.valsPreOrder, extracValuesPreOrder(root); !cmp.Equal(want, got) {
				printTreePreOrder(t, root)
				t.Fatalf("want= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got))
			}
		})
	}
}

func TestExtractValuesPostOrder(t *testing.T) {
	var testCases = map[string]struct {
		vals          []containers.Value
		valsPostOrder []containers.Value
	}{
		// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder
		"geeksForGeeks": {
			vals:          []containers.Value{1, 2, 3, 4, 5},
			valsPostOrder: []containers.Value{4, 5, 2, 3, 1},
		},
		// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/992
		"leetCode": {
			vals:          []containers.Value{"F", "B", "G", "A", "D", nil, "I", nil, nil, "C", "E", nil, nil, "H"},
			valsPostOrder: []containers.Value{"A", "C", "E", "D", "B", "H", "I", "G", "F"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			root, err := NewFromLeetCodeOrder(tc.vals...)
			if err != nil {
				t.Fatalf("need a valid btree to test")
			}

			if want, got := tc.valsPostOrder, extracValuesPostOrder(root); !cmp.Equal(want, got) {
				printTreePreOrder(t, root)
				t.Fatalf("want= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got))
			}
		})
	}
}

func TestExtractValuesInOrder(t *testing.T) {
	var testCases = map[string]struct {
		vals        []containers.Value
		valsInOrder []containers.Value
	}{
		// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder
		"geeksForGeeks": {
			vals:        []containers.Value{1, 2, 3, 4, 5},
			valsInOrder: []containers.Value{4, 2, 5, 1, 3},
		},
		// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/992
		"leetCode": {
			vals:        []containers.Value{"F", "B", "G", "A", "D", nil, "I", nil, nil, "C", "E", nil, nil, "H"},
			valsInOrder: []containers.Value{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			root, err := NewFromLeetCodeOrder(tc.vals...)
			if err != nil {
				t.Fatalf("need a valid btree to test")
			}

			if want, got := tc.valsInOrder, extracValuesInOrder(root); !cmp.Equal(want, got) {
				printTreePreOrder(t, root)
				t.Fatalf("want= %v, got= %v, diff= %v", want, got, cmp.Diff(want, got))
			}
		})
	}
}

func printTreePreOrder(t *testing.T, root *TreeNode) {
	walkPreOrder(root, func(node *TreeNode) {
		if node.Value == nil {
			t.Logf("(@ %12p, V= %2v, P= %12p, L= %12p, R= %12p)", node, "-", node.Parent, node.Left, node.Right)
			return
		}

		t.Logf("(@ %12p, V= %2v, P= %12p, L= %12p, R= %12p)", node, node.Value, node.Parent, node.Left, node.Right)
	})
}
