package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		p_prev := head
		p := head
		q_prev := head
		q := q_prev.Next

		for q != nil {
			q_next := q.Next
			if q.Val < p.Val {
				head = moveAhead(head, p, p_prev, q, q_prev)
				p_prev = q
			} else {
				q_prev = q_prev.Next
			}
			q = q_next
		}

		p_next := p.Next
		p.Next = nil
		head := sortList(head)
		sortList(head)

		rh := sortList(p_next)
		p.Next = rh

		return head
	}
}

func moveAhead(head *ListNode, p *ListNode, p_prev *ListNode, q *ListNode, q_prev *ListNode) *ListNode {
	q_prev.Next = q.Next
	q.Next = p
	if p_prev != p {
		p_prev.Next = q
		return head
	} else {
		return q
	}
}
