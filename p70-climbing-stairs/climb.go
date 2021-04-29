package main

var cache = map[int]int{1: 1, 2: 2}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return climbStairs(n)
}
