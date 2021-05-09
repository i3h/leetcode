package main

import "fmt"

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(nums)
	nums = sortArrayByParityII(nums)
	fmt.Println(nums)
}
