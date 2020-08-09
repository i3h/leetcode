package main

func search(nums []int, target int) int {
	l := len(nums)
	a := 0
	c := l - 1
	b := (a + c) / 2
	for {
		if target < nums[b] {
			c = b
		} else {
			a = b
		}
		b = (a + c) / 2
		if a == b || b == c {
			if target == nums[a] {
				//fmt.Println(a, b, c)
				return a
			} else if target == nums[c] {
				//fmt.Println(a, b, c)
				return c
			} else {
				//fmt.Println(a, b, c)
				return -1
			}
		}
	}
}
