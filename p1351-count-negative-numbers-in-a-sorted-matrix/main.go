package main

import "fmt"

func main() {
	//grid := [][]int{[]int{4, 3, 2, -1}, []int{3, 2, 1, -1}, []int{1, 1, -1, -2}, []int{-1, -1, -2, -3}}
	grid := [][]int{[]int{5, 1, 0}, []int{-5, -5, -5}}
	printGrid(grid)
	fmt.Println(countNegatives(grid))
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%2d ", grid[i][j])
		}
		fmt.Println()
	}
}
