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

// walkPreOrder executes fn() on nodes using pre-order.
func walkPreOrder(node *TreeNode, fn func(*TreeNode)) {
	if node == nil {
		return
	}

	fn(node)
	walkPreOrder(node.Left, fn)
	walkPreOrder(node.Right, fn)
}

// walkPostOrder executes fn() on nodes using post-order.
func walkPostOrder(node *TreeNode, fn func(*TreeNode)) {
	if node == nil {
		return
	}

	walkPostOrder(node.Left, fn)
	walkPostOrder(node.Right, fn)
	fn(node)
}

// walkInOrder executes fn() on nodes using in-order.
func walkInOrder(node *TreeNode, fn func(*TreeNode)) {
	if node == nil {
		return
	}

	walkInOrder(node.Left, fn)
	fn(node)
	walkInOrder(node.Right, fn)
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
	root.Left = nodeFromLeetCodeOrder(vals, 1, &root)
	root.Right = nodeFromLeetCodeOrder(vals, 2, &root)

	return &root, nil
}

func nodeFromLeetCodeOrder(vals []containers.Value, index int, parent *TreeNode) *TreeNode {
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
	node.Left = nodeFromLeetCodeOrder(vals, index*2+1, &node)
	node.Right = nodeFromLeetCodeOrder(vals, index*2+2, &node)

	return &node
}

func extracValuesPreOrder(root *TreeNode) []containers.Value {
	return extractValues(root, walkPreOrder)
}

func extracValuesPostOrder(root *TreeNode) []containers.Value {
	return extractValues(root, walkPostOrder)
}

func extracValuesInOrder(root *TreeNode) []containers.Value {
	return extractValues(root, walkInOrder)
}

func extractValues(root *TreeNode, traversalFn func(*TreeNode, func(*TreeNode))) []containers.Value {
	var vals []containers.Value

	appendFn := func(node *TreeNode) {
		vals = append(vals, node.Value)
	}
	traversalFn(root, appendFn)

	return vals
}
