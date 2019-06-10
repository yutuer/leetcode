package _617

import "testing"

func TestMergeTrees1(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}

	tree := mergeTrees(t1, t2)
	tree.pt()
}

func TestMergeTrees2(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}

	tree := mergeTrees(t1, t2)
	tree.pt()
}
