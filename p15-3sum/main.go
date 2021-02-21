package main

import "fmt"

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0, 0}
	fmt.Println("Input: ", nums)

	set := threeSum(nums)
	fmt.Println("Output: ", set)
}
