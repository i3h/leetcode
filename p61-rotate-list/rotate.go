package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 1
	p := head
	for p.Next != nil {
		n++
		p = p.Next
	}
	k = k % n

	p.Next = head

	for i := 0; i < n-k; i++ {
		if i == n-k-1 {
			q := head.Next
			head.Next = nil
			head = q
		} else {
			head = head.Next
		}
	}

	return head
}
