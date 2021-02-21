package main

import "fmt"

func main() {
	n := 2
	nodes := make([]ListNode, n)
	for i := 0; i < n; i++ {
		if i < n-1 {
			nodes[i] = ListNode{
				Val:  i + 1,
				Next: &nodes[i+1],
			}
		} else {
			nodes[i] = ListNode{
				Val:  i + 1,
				Next: nil,
			}
		}
	}
	//fmt.Println(nodes)
	head := &nodes[0]
	head = removeNthFromEnd(head, 2)
	printList(head)
}

func printList(head *ListNode) {
	p := head
	fmt.Println("==============")
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
	fmt.Println("==============")
}
