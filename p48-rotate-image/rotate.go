package main

import "fmt"

func rotate(matrix [][]int) {
	/*
	 * (x,y)
	 * (y,l-x)
	 * (l-x,l-y)
	 * (l-y,x)
	 * l is the max value of x,y
	 */
	l := len(matrix[0]) - 1
	fmt.Println("l: ", l)
	c := len(matrix[0]) / 2
	fmt.Println("circles: ", c)
	y := 0
	for x := 0; x < l; x++ {
		fmt.Printf("swap: (%d, %d) -> (%d, %d)\n", x, y, y, l-x)
		matrix[x][y], matrix[y][l-x] = matrix[y][l-x], matrix[x][y]
		printMatrix(matrix)

		fmt.Printf("swap: (%d, %d) -> (%d, %d)\n", x, y, l-x, l-y)
		matrix[x][y], matrix[l-x][l-y] = matrix[l-x][l-y], matrix[x][y]
		printMatrix(matrix)

		fmt.Printf("swap: (%d, %d) -> (%d, %d)\n", x, y, l-y, x)
		matrix[x][y], matrix[l-y][x] = matrix[l-y][x], matrix[x][y]
		printMatrix(matrix)
	}
}
