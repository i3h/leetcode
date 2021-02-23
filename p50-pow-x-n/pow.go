package main

func myPow(x float64, n int) float64 {
	a := x
	if x < 0 {
		a = -a
	}
	if a > 1e6 || a < 1e-6 {
		return x
	}
	if x == 1 {
		return 1
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}

	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n == -1 {
		return 1 / x
	} else {
		if n%2 == 0 {
			return myPow(x, n/2) * myPow(x, n/2)
		} else {
			if n > 0 {
				return myPow(x, n/2) * myPow(x, n/2+1)
			} else {
				return myPow(x, n/2) * myPow(x, n/2-1)
			}
		}
	}
}
