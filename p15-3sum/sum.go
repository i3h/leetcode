package main

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	var set, tmp_set [][]int
	sort.Ints(nums)
	//fmt.Println("Sort: ", nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		var s []int
		for j := i + 1; j < len(nums); j++ {
			if j != i {
				s = append(s, nums[j])
			}
		}
		tmp_set = append(tmp_set, twoSum(s, -nums[i])...)
	}

	set = cleanSet(tmp_set)

	return set
}

func cleanSet(set [][]int) [][]int {
	var new_set [][]int

	m := make(map[string][]int)
	for _, t := range set {
		fmt.Println(t)
		s := strconv.Itoa(t[0]) + "+" + strconv.Itoa(t[1]) + "+" + strconv.Itoa(t[2])
		m[s] = t
		fmt.Println(s)
	}
	fmt.Println(m)

	for _, v := range m {
		new_set = append(new_set, v)
	}

	return new_set
}

func twoSum(nums []int, target int) [][]int {
	var set [][]int

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		com := target - nums[i]
		if val, ok := m[com]; ok && i != val {
			set = append(set, []int{-target, nums[i], nums[val]})
			//break
		}
		m[nums[i]] = i
	}
	for i := 0; i < len(set); i++ {
		sort.Ints(set[i])
	}

	return set
}
