package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := createList(1)
	l2 := createList(9)
	printList(l1)
	fmt.Println()
	printList(l2)
	fmt.Println()
	l3 := addTwoNumbers(l1, l2)
	printList(l3)
	fmt.Println()
}

func createList(nums ...int) *ListNode {
	var l, p *ListNode

	for i, n := range nums {
		if i == 0 {
			p = &ListNode{Val: n, Next: nil}
			l = p
		} else {
			p.Next = &ListNode{Val: n, Next: nil}
			p = p.Next
		}
	}

	return l
}

func printList(p *ListNode) {
	for {
		if p.Next == nil {
			fmt.Println(p.Val)
			return
		} else {
			fmt.Println(p.Val)
			p = p.Next
		}
	}
}
