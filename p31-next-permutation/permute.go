package main

import (
	"sort"
)

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
