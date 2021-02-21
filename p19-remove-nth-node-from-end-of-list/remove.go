package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	q := head
	c := 0

	for i := 0; i < n+1; i++ {
		if q == nil {
			if c == 1 {
				return nil
			} else {
				return head.Next
			}
		}
		q = q.Next
		c++
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next

	return head
}
