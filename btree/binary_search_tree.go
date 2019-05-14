package btree

import (
	"github.com/bitsgofer/containers"
)

// IsBST validates whether a btree is a BST.
func IsBST(root *TreeNode, less func(a, b containers.Value) bool, minusInf, plusInf containers.Value) bool {
	// traverse down the tree inorder, check that values in each subtree lies is in (min, max).
	// start at root with (-inf, +inf)

	return isBST(root, less, minusInf, plusInf)
}

func isBST(root *TreeNode, less func(a, b containers.Value) bool, min, max containers.Value) bool {
	if root == nil {
		return true
	}

	var leftNotValid, rightNotValid bool
	if root.Left != nil {
		leftLessThanRoot := less(root.Left.Value, root.Value)
		leftLessThanMax := true // left < root && root < max
		leftGreaterThanMin := less(min, root.Left.Value)
		leftIsBST := isBST(root.Left, less, min, root.Value)

		leftNotValid = !leftLessThanRoot || !leftLessThanMax || !leftGreaterThanMin || !leftIsBST
	}
	if root.Right != nil {
		rightGreaterThanRoot := less(root.Value, root.Right.Value)
		rightLessThanMax := less(root.Right.Value, max)
		rightGreaterThanMin := true // min < root && root < right
		rightIsBST := isBST(root.Right, less, root.Value, max)

		rightNotValid = !rightGreaterThanRoot || !rightLessThanMax || !rightGreaterThanMin || !rightIsBST
	}
	return !leftNotValid && !rightNotValid
}
