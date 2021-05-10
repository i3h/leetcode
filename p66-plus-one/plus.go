package main

func plusOne(digits []int) []int {
	overflow := false

	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] < 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
			if i == 0 {
				overflow = true
			}
		}
	}

	if overflow {
		digits = append([]int{1}, digits...)
	}

	return digits
}
