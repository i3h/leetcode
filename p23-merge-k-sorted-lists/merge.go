package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	//fmt.Println("head: ", head)
	q := head
	pList := make([]*ListNode, len(lists))
	for i := 0; i < len(pList); i++ {
		pList[i] = lists[i]
		//fmt.Println(pList[i].Val)
	}
	for !allNil(pList) {
		//fmt.Println("head: ", head)
		//fmt.Println("q   : ", q)
		p := min(pList)
		/*
			fmt.Println("pList:   ", pList[0], pList[1], pList[2])
			fmt.Println("p:       ", p)
			fmt.Println("pList_p: ", pList[p])
		*/
		q.Next = &ListNode{
			Val:  pList[p].Val,
			Next: nil,
		}
		q = q.Next
		//fmt.Println("HEAD BEGIN")
		//printList(head)
		//fmt.Println("HEAD END")
		pList[p] = pList[p].Next
	}

	return head.Next
}

func min(lists []*ListNode) int {
	//fmt.Println("lists:   ", lists[0], lists[1], lists[2])
	var p, m int
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			p = i
			m = lists[i].Val
			break
		}
	}
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		} else {
			if m > lists[i].Val {
				p = i
				m = lists[i].Val
			}
		}
	}
	return p
}

func allNil(lists []*ListNode) bool {
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			return false
		}
	}
	return true
}
