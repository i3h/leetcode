package main

import "fmt"

func main() {
	nums := []int{3, 2, 0, -4}
	head := setupList(nums)
	//printList(head)
	fmt.Println(detectCycle(head))
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

	p.Next = head.Next

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
