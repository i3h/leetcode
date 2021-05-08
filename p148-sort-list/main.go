package main

import "fmt"

func main() {
	//nums := []int{-1, 5, 3, 4, 0}
	nums := []int{4, 19, 14, 5, -3, 1, 8, 5, 11, 15}
	//nums := []int{-1, 5, 4, 0}
	//nums := []int{3, 4, 0, 5}
	//nums := []int{3, 3}
	head := setupList(nums)
	printList(head)
	//sortList(head)
	printList(sortList(head))
}

func setupList(nums []int) *ListNode {
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	p := head

	for i := 1; i < len(nums); i++ {
		p.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		p = p.Next
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
