package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := p1.Next
	l1 := 1
	l2 := 2
	a := 0
	b := 0
	x := 0
	for p2 != nil {
		if p1 != p2 {
			p1 = p1.Next
			l1++
			p2 = p2.Next
			l2++
			if p2 != nil {
				p2 = p2.Next
				l2++
			}
		} else {
			p0 := p2.Next
			ab := 2
			for p0 != p2 {
				p0 = p0.Next
				ab++
			}
			//fmt.Println("ab = ", ab)
			b = l2 - l1
			a = ab - b
			x = (l1 + l2 - 2*a - b) / 2
			//fmt.Printf("l1 = %d l2 = %d a = %d b = %d x = %d\n", l1, l2, a, b, x)
			px := head
			for i := 0; i < x-1; i++ {
				px = px.Next
			}
			return px
		}
	}
	return nil
}
