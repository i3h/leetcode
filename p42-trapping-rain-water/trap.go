package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	c := 0
	max := findMax(height)
	//fmt.Println("max: ", max)

	for i := max - 1; i >= 0; i-- {
		tmp := make([]int, len(height))
		copy(tmp, height)
		for j := 0; j < len(tmp); j++ {
			tmp[j] -= i
			if tmp[j] < 0 {
				tmp[j] = 0
			}
		}
		//fmt.Println("tmp: ", tmp)
		c += countTrap(tmp)
	}

	return c
}

func countTrap(height []int) int {
	c := 0
	p, q := 0, 0
	for i := 0; i < len(height); i++ {
		if height[i] != 0 {
			p = i
			break
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] != 0 {
			q = i
			break
		}
	}
	//fmt.Println("p: ", p)
	//fmt.Println("q: ", q)
	for i := p; i < q; i++ {
		if height[i] == 0 {
			c++
		}
	}
	//fmt.Println("count: ", c)
	return c
}

func findMax(height []int) int {
	max := height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	return max
}
