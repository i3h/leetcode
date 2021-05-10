package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(h1 *TreeNode, h2 *TreeNode) bool {
	if h1 == nil && h2 == nil {
		return true
	}
	if h1 == nil || h2 == nil {
		return false
	}
	if h1.Val != h2.Val {
		return false
	}
	return isMirror(h1.Left, h2.Right) && isMirror(h1.Right, h2.Left)
}
