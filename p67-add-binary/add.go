package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	l := len(a)
	a1 := make([]int, l)
	a2 := make([]int, l)
	for i := 0; i < len(a); i++ {
		a1[i], _ = strconv.Atoi(string(a[len(a)-i-1]))
		if i < len(b) {
			a2[i], _ = strconv.Atoi(string(b[len(b)-i-1]))
		} else {
			a2[i] = 0
		}
	}

	r := 0
	a3 := make([]int, l+1)
	for i := 0; i < l; i++ {
		s := a1[i] + a2[i] + r
		a3[i] = s % 2
		r = s / 2
	}
	if r == 1 {
		a3[l] = 1
	}

	found := false
	for i := len(a3) - 1; i >= 0; i-- {
		if a3[i] != 0 {
			a3 = a3[:i+1]
			found = true
			break
		}
	}

	var c string
	for i := len(a3) - 1; i >= 0; i-- {
		c += strconv.Itoa(a3[i])
	}
	if !found {
		c = "0"
	}

	return c
}
