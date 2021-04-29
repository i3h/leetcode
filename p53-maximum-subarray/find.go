package main

import "fmt"

var MAX int
var TAB [][]int

func maxSubArray(nums []int) int {
	fmt.Println(MAX)
	fmt.Println(nums)
	if len(nums) == 0 {
		return MAX
	} else if len(nums) == 1 {
		if nums[0] > MAX {
			return nums[0]
		} else {
			return MAX
		}
	} else {
		compare(nums)
		compare(nums[:len(nums)-2])
		compare(nums[1 : len(nums)-1])
		return maxSubArray(nums[1 : len(nums)-2])
	}
}

func compare(nums []int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum > MAX {
		MAX = sum
	}
}
