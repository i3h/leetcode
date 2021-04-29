package main

/*
 * T_0 = 0
 * T_1 = 1
 * T_2 = 1
 * T_n+3 = T_n + T_n+1 + T_n+2 for n >= 0
 */

var cache = map[int]int{0: 0, 1: 1, 2: 1}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	cache[n] = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	return cache[n]
}
