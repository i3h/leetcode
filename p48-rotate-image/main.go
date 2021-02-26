package main

import "fmt"

func main() {
	matrix := [][]int{}
	/*
		matrix = append(matrix, []int{1, 2, 3})
		matrix = append(matrix, []int{4, 5, 6})
		matrix = append(matrix, []int{7, 8, 9})
	*/
	matrix = append(matrix, []int{5, 1, 9, 11})
	matrix = append(matrix, []int{2, 4, 8, 10})
	matrix = append(matrix, []int{13, 3, 6, 7})
	matrix = append(matrix, []int{15, 14, 12, 16})
	printMatrix(matrix)
	rotate(matrix)
	printMatrix(matrix)
}

func printMatrix(m [][]int) {
	fmt.Println("========================")
	l := len(m[0])
	for i := 0; i < l; i++ {
		fmt.Println(m[i])
	}
	fmt.Println("========================")
}
