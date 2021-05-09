package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println("nums1: ", nums1)
	fmt.Println("nums2: ", nums2)
	merge(nums1, m, nums2, n)
	fmt.Println()
	fmt.Println("nums1: ", nums1)
}
