package main

func countNegatives(grid [][]int) int {
	c := 0
	p := len(grid[0]) - 1
	for i := 0; i < len(grid); i++ {
		//fmt.Println("line: ", grid[i])
		for j := p; j >= 0; j-- {
			if grid[i][j] >= 0 {
				//fmt.Println(grid[i][j])
				p = j
				c += p + 1
				break
			}
		}
		//fmt.Printf("c = %d p = %d\n", c, p)
	}

	return len(grid)*len(grid[0]) - c
}
