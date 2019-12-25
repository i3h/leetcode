package main

func hasRepeat(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			return true
		}
		m[s[i]] = 1
	}
	return false
}

func getMax(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
	return n1
}

func lengthOfLongestSubstring(s string) int {
	l := 0
	p1 := 0
	p2 := 0

	for i := 0; i < len(s)+1; i++ {
		if p2 > len(s) {
			break
		}
		if !hasRepeat(s[p1:p2]) {
			l = getMax(l, p2-p1)
			p2++
		} else {
			p1++
			p2++
		}
	}

	return l
}
