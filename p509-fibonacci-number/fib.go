package main

/*
 * f(0) = 0
 * f(1) = 1
 * f(n) = f(n-1) + f(n-2), for n > 1.
 */

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
