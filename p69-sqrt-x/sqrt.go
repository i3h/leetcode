package main

import "fmt"

/*
 * Newton Method
 * x_n+1 = (x_n^2 + a) / (2 * X_n)
 */

func mySqrt(x int) int {
	var r float64
	r = float64(x) / 2

	for i := 0; i < 10; i++ {
		r = (r*r + float64(x)) / 2 / r
	}

	fmt.Println("r = ", r)
	return int(r)
}
