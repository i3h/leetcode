package main

import "fmt"

func main() {
	//nums := []int{1, 5, 1, 1, 6, 4}
	nums := []int{1, 3, 2, 2, 3, 1}
	fmt.Println(nums)
	wiggleSort(nums)
	fmt.Println(nums)
}
