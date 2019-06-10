package _965

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	return compare(root, root.Val)
}

func compare(node *TreeNode, i int) bool {
	if node == nil {
		return true
	}

	if node.Val != i {
		return false
	}

	if (compare(node.Left, i)) {
		return compare(node.Right, i)
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
