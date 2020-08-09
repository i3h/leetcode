package main

import "fmt"

func main() {
	nums := []int{2, 5}
	target := 5
	i := search(nums, target)
	fmt.Println(i)
}
