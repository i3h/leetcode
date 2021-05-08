package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next

	for q != nil {
		if p.Val == q.Val {
			p.Next = q.Next
			q = q.Next
		} else {
			p = p.Next
			q = q.Next
		}
	}

	return head
}
