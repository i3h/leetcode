package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	p, q := l1, l2
	if p.Val > q.Val {
		p, q = q, p
	}
	h := p

	for p.Next != nil && q.Next != nil {
		if q.Val < p.Next.Val {
			t := p.Next
			p.Next = q
			p = p.Next
			q = q.Next
			p.Next = t
		} else {
			p = p.Next
		}
	}

	if p.Next == nil {
		p.Next = q
	} else {
		for p.Next != nil {
			if p.Next.Val < q.Val {
				p = p.Next
			} else {
				q.Next = p.Next
				p.Next = q
				break
			}
		}

		if q != nil {
			p.Next = q
		}
	}

	return h
}
