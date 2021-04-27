package main

func rotate(matrix [][]int) {
	/*
	 * (x,y)
	 * (y,l-x)
	 * (l-x,l-y)
	 * (l-y,x)
	 * l is the max value of x,y
	 */
	l := len(matrix[0]) - 1
	c := len(matrix[0]) / 2
	for i := 0; i < c; i++ {
		y := i
		for x := i; x < l-i; x++ {
			matrix[x][y], matrix[y][l-x] = matrix[y][l-x], matrix[x][y]
			matrix[x][y], matrix[l-x][l-y] = matrix[l-x][l-y], matrix[x][y]
			matrix[x][y], matrix[l-y][x] = matrix[l-y][x], matrix[x][y]
		}
	}
}
