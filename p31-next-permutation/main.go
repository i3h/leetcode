package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3}
	//nums := []int{3, 2, 1}
	//nums := []int{3, 4, 2, 1}
	//nums := []int{4, 2, 0, 2, 3, 2, 0}
	//nums := []int{1, 3, 2}
	nums := []int{2, 2, 0, 4, 3, 1}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
}
