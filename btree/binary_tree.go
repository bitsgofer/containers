package btree

import (
	"github.com/pkg/errors"

	"github.com/bitsgofer/containers"
)

// TreeNode is a node in the binary tree.
type TreeNode struct {
	Value  containers.Value
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

// walkPreorder traverse a tree inorder, executing fn() on each node.
func walkPreorder(node *TreeNode, fn func(*TreeNode)) {
	if node == nil {
		return
	}

	fn(node)
	walkPreorder(node.Left, fn)
	walkPreorder(node.Right, fn)
}

// NewFromLeetCodeOrder construct a btree from a list of values (LeetCode test cases).
// The values are given as if we are traversing a full binary tree with BFS, with nil representing empty nodes.
func NewFromLeetCodeOrder(vals ...containers.Value) (*TreeNode, error) {
	if len(vals) == 0 {
		return nil, errors.New("no values")
	}

	root := TreeNode{
		Value: containers.Value(vals[0]),
	}
	root.Left = nodeFromBFS(vals, 1, &root)
	root.Right = nodeFromBFS(vals, 2, &root)

	return &root, nil
}

func nodeFromBFS(vals []containers.Value, index int, parent *TreeNode) *TreeNode {
	if index >= len(vals) { // no such element
		return nil
	}

	val := vals[index]
	if val == nil {
		// no node => also won't consider indices that would have been children of this node
		return nil
	}

	node := TreeNode{
		Value:  containers.Value(vals[index]),
		Parent: parent,
	}
	node.Left = nodeFromBFS(vals, index*2+1, &node)
	node.Right = nodeFromBFS(vals, index*2+2, &node)

	return &node
}
