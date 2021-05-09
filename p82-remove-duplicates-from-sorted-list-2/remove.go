package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Val == head.Next.Val {
			return nil
		} else {
			return head
		}
	}

	i := head
	j := head.Next
	k := head.Next.Next
	skip := false

	//if i.Val == j.Val {
	for i.Val == j.Val {
		//fmt.Printf("i = %d j = %d k = %d\n", i.Val, j.Val, k.Val)
		for j != nil && j.Val == i.Val {
			j = j.Next
		}
		if j == nil {
			return nil
		}
		//fmt.Printf("i = %d j = %d k = %d\n", i.Val, j.Val, k.Val)
		head = j
		i = j
		j = i.Next
		if j != nil {
			k = j.Next
		} else {
			return head
		}
		//fmt.Printf("i = %d j = %d\n", i.Val, j.Val)
	}

	for k != nil {
		//fmt.Printf("i = %d j = %d k = %d\n", i.Val, j.Val, k.Val)
		if j.Val == k.Val {
			k = k.Next
			skip = true
			if k == nil {
				i.Next = k
			}
		} else {
			if !skip {
				i = i.Next
				j = j.Next
				k = k.Next
			} else {
				i.Next = k
				j = k
				k = k.Next
				skip = false
			}
		}
	}

	return head
}
