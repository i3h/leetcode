package main

import "fmt"

func main() {
	//nums := []int{5, 2, 3, 1}
	//nums := []int{5, 1, 1, 2, 0, 0}
	nums := []int{-1, 2, -8, -10}
	fmt.Println(nums)
	nums = sortArray(nums)
	fmt.Println(nums)
}
