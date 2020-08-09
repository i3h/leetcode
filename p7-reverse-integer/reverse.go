package main

import (
	"math"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	digits := [32]int{}
	for i := 0; i < 32; i++ {
		b := int(math.Pow(10, float64(i)))
		if x < b {
			break
		}
		digits[i] = (x%(b*10) - x%b) / b
	}
	//fmt.Println(digits)
	var y int
	var j int
	isValid := false
	for i := 31; i > -1; i-- {
		if digits[i] == 0 && !isValid {
			continue
		} else {
			isValid = true
			y += digits[i] * int(math.Pow(10, float64(j)))
			j++
		}
	}
	if float64(y) > math.Pow(2, 31)-1 {
		return 0
	}
	if isNegative {
		return -y
	} else {
		return y
	}
}
