package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println("nums: ", nums)
	n := findDuplicate(nums)
	fmt.Println(n)
}
