package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	//target := 5
	target := 7
	fmt.Println(nums)
	fmt.Println(searchInsert(nums, target))
}
