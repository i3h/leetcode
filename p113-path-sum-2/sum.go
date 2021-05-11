package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	m := make(map[*TreeNode][]int)
	searchPathSum(root, targetSum, m, nil)
	fmt.Println(m)
	return nil
}

func searchPathSum(root *TreeNode, targetSum int, m map[*TreeNode][]int, leaf *TreeNode) (bool, *TreeNode) {
	if root == nil {
		return false, nil
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		m[root] = append(m[root], root.Val)
		return true, root
	} else {
		l, l_leaf := searchPathSum(root.Left, targetSum-root.Val, m, nil)
		r, r_leaf := searchPathSum(root.Right, targetSum-root.Val, m, nil)
		if !l {
			root.Left = nil
		} else {
			m[l_leaf] = append(m[l_leaf], l_leaf.Val)
			return true, l_leaf
		}
		if !r {
			root.Right = nil
		} else {
			m[r_leaf] = append(m[r_leaf], r_leaf.Val)
			return true, r_leaf
		}
		return false, nil
	}
}
