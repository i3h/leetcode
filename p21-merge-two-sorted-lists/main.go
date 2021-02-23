package main

import "fmt"

func main() {
	n := 3
	m := 1
	l1 := make([]ListNode, n)
	l2 := make([]ListNode, m)
	for i := 0; i < n-1; i++ {
		l1[i].Next = &l1[i+1]
	}
	for i := 0; i < m-1; i++ {
		l2[i].Next = &l2[i+1]
	}

	l1[0].Val = 1
	l1[1].Val = 2
	l1[2].Val = 4

	l2[0].Val = 3

	printList(&l1[0])
	printList(&l2[0])

	head := mergeTwoLists(&l1[0], &l2[0])
	printList(head)
}

func printList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}
