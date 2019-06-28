package _700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BST是二叉搜索树, 已经有序了
//func searchBST(root *TreeNode, val int) *TreeNode {
//	// 深度优先搜索
//	if root == nil {
//		return nil
//	}
//
//	if root.Val == val {
//		return root
//	}
//
//	if val > root.Val {
//		return searchBST(root.Right, val)
//	} else {
//		return searchBST(root.Left, val)
//	}
//}

func searchBST(root *TreeNode, val int) *TreeNode {
	for n := root; n != nil; {
		if n.Val == val {
			return n
		} else if val > n.Val {
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return nil
}
