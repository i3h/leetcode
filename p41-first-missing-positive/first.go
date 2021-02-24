package main

import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	} else if len(nums) == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	sort.Ints(nums)
	//fmt.Println(nums)

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	//fmt.Println(m)

	middle := findMiddle(nums)
	//fmt.Println("middle: ", middle)
	p := m[middle]
	//fmt.Println("position: ", p)

	var min int
	if nums[p] > 0 {
		min = nums[p]
	} else {
		if p > len(nums)-2 {
			return 1
		} else {
			min = nums[p+1]
		}
	}
	//fmt.Println("min: ", min)

	if min > 1 {
		return 1
	} else {
		i := 2
		for {
			if _, ok := m[i]; !ok {
				return i
			}
			i++
		}
	}
}

func findMiddle(nums []int) int {
	l := len(nums)
	//fmt.Println("nums: ", nums)
	if l == 1 {
		return nums[0]
	}
	if nums[l/2] > 0 {
		return findMiddle(nums[:l/2])
	} else {
		return findMiddle(nums[l/2:])
	}
}
