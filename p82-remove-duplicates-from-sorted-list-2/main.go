package main

import "fmt"

func main() {
	//nums := []int{1, 1, 2, 3, 3}
	//nums := []int{3, 3}
	//nums := []int{1, 2, 3, 3, 4, 4, 5}
	nums := []int{1, 1, 2, 2}
	head := setupList(nums)
	printList(head)
	head = deleteDuplicates(head)
	printList(head)
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
