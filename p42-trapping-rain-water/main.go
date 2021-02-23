package main

import "fmt"

func main() {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("height: ", h)
	t := trap(h)
	fmt.Println(t)
}
