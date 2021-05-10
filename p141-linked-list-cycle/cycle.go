package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	p := head
	for p != nil {
		if m[p] == 0 {
			m[p] = 1
			p = p.Next
		} else {
			return true
		}
	}
	return false
}
