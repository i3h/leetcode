package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{Val: 0, Next: nil}
	p := l
	v1 := 0
	v2 := 0
	quo := 0
	p1 := l1
	p2 := l2

	for {
		if p1 == nil && p2 == nil {
			if quo != 0 {
				s := &ListNode{Val: quo, Next: nil}
				p.Next = s
			}
			break
		}
		if p1 != nil {
			v1 = p1.Val
			p1 = p1.Next
		} else {
			v1 = 0
		}
		if p2 != nil {
			v2 = p2.Val
			p2 = p2.Next
		} else {
			v2 = 0
		}
		s := &ListNode{Val: (v1 + v2 + quo) % 10, Next: nil}
		quo = (v1 + v2 + quo) / 10
		p.Next = s
		p = s
	}

	return l.Next
}
