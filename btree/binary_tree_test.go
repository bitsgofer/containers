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
				if want, got := tc.valsInOrder, extracValuesPreorder(root); !cmp.Equal(want, got) {
					t.Fatalf("values inorder: want %v, got %v, diff= %v", want, got, cmp.Diff(want, got))
				}
			}
		})
	}
}

func printTreePreorder(t *testing.T, root *TreeNode) {
	walkPreorder(root, func(node *TreeNode) {
		if node.Value == nil {
			t.Logf("(@ %12p, V= %2v, P= %12p, L= %12p, R= %12p)", node, "-", node.Parent, node.Left, node.Right)
			return
		}

		t.Logf("(@ %12p, V= %2v, P= %12p, L= %12p, R= %12p)", node, node.Value, node.Parent, node.Left, node.Right)
	})
}

func extracValuesPreorder(root *TreeNode) []containers.Value {
	var vals []containers.Value

	walkPreorder(root, func(node *TreeNode) {
		vals = append(vals, node.Value)
	})

	return vals
}
