package main

type Node struct {
	Val      int
	Children []*Node
}

var nodes = []int{}

func preorder(root *Node) []int {
	nodes = append(nodes, root.Val)
	for _, node := range root.Children {
		preorder(node)
	}
	return nodes
}
