package main

import "fmt"

func main() {
	nums := [][]int{[]int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}}
	lists := setupList(nums)
	printLists(lists)
	m := mergeKLists(lists)
	printList(m)
}

func setupList(nums [][]int) []*ListNode {
	var lists []*ListNode
	for i := 0; i < len(nums); i++ {
		head := &ListNode{
			Val:  nums[i][0],
			Next: nil,
		}
		p := head

		for j := 1; j < len(nums[i]); j++ {
			p.Next = &ListNode{
				Val:  nums[i][j],
				Next: nil,
			}
			p = p.Next
		}

		lists = append(lists, head)
	}
	return lists
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func printLists(lists []*ListNode) {
	for i := 0; i < len(lists); i++ {
		head := lists[i]
		for head != nil {
			fmt.Printf("%d ", head.Val)
			head = head.Next
		}
		fmt.Println()
	}
}
