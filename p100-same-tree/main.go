package main

import "fmt"

func main() {
	// Create p
	p := &TreeNode{
		Val: 1,
	}
	p1 := &TreeNode{
		Val: 2,
	}
	p2 := &TreeNode{
		Val: 3,
	}
	p.Left = p1
	p.Right = p2
	// Create q
	q := &TreeNode{
		Val: 1,
	}
	q1 := &TreeNode{
		Val: 3,
	}
	q2 := &TreeNode{
		Val: 2,
	}
	q.Left = q1
	q.Right = q2
	fmt.Println(isSameTree(p, q))
}
