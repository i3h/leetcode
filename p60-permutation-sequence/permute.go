package main

import (
	"sort"
)

func getPermutation(n int, k int) string {
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = i + 1 + 48
	}

	for i := 0; i < k-1; i++ {
		nextPermutation(nums)
	}

	set := make([]rune, n)
	for i := 0; i < n; i++ {
		set[i] = rune(nums[i])
	}

	return string(set)
}

func nextPermutation(nums []int) {
	found := false

out:
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			found = true
			for j := len(nums) - 1; j > i-1; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					sort.Ints(nums[i:])
					break out
				}
			}
		}
	}

	if !found {
		// reverse
		for i := 0; i < len(nums)/2; i++ {
			j := len(nums) - i - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
