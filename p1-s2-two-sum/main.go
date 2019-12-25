package main

import "fmt"

func main() {
	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 3}
	target := 6
	indexs := twoSum(nums, target)
	fmt.Println(indexs)
}
