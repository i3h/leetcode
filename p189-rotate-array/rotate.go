package main

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}
	r := l - k%l
	list := make([]int, l)
	copy(list, nums)
	for i := 0; i < l-r; i++ {
		nums[i] = list[i+r]
	}
	for i := l - r; i < l; i++ {
		nums[i] = list[i+r-l]
	}
}
