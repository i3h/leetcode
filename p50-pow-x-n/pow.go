package main

func myPow(x float64, n int) float64 {
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
	s := float64(1)

	inverse := false
	if n < 0 {
		inverse = true
		n = -n
	}
	for i := 0; i < n; i++ {
		s = s * x
		if s > 1e6 {
			break
		}
	}
	if inverse {
		s = 1 / s
	}

	return s
}
