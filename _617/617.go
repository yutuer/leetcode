package _617

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return merge(t1, t2)
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	} else {
		t := &TreeNode{}
		t.Val = t1.Val + t2.Val
		t.Left = merge(t1.Left, t2.Left)
		t.Right = merge(t1.Right, t2.Right)
		return t
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) pt() {
	p(t)
}

func p(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Val)
		p(t.Left)
		p(t.Right)
	}
}
