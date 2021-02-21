package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var set [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		set = append(set, twoSum(nums[i+1:], -nums[i])...)
	}

	return set
}

func twoSum(nums []int, target int) [][]int {
	var set [][]int

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		com := target - nums[i]
		if val, ok := m[com]; ok && i != val {
			set = append(set, []int{nums[i], nums[val]})
		}
		m[nums[i]] = i
	}
	for i := 0; i < len(set); i++ {
		sort.Ints(set[i])
	}

	ms := make(map[int][]int)
	for _, s := range set {
		ms[s[0]] = s
	}
	set = [][]int{}
	for _, v := range ms {
		v = append(v, -target)
		set = append(set, v)
	}

	return set
}
