package main

/*
 * f(0) = 0
 * f(1) = 1
 * f(n) = f(n-1) + f(n-2), for n > 1.
 */

var CACHE = map[int]int{0: 0, 1: 1}

func fib(n int) int {
	if n < 2 {
		return n
	}
	if _, ok := CACHE[n]; ok {
		return CACHE[n]
	}
	CACHE[n] = fib(n-1) + fib(n-2)
	return fib(n)
}
