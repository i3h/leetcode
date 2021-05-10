package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + maxDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + maxDepth(root.Left)
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l < r {
		return 1 + r
	} else {
		return 1 + l
	}
}
