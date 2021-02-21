package main

import "sort"

func permuteUnique(nums []int) [][]int {
	var set [][]int
	sort.Ints(nums)
	if len(nums) == 1 {
		return [][]int{nums}
	} else {
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			s := []int{}
			k := []int{}
			s = append(s, nums[:i]...)
			s = append(s, nums[i+1:]...)
			k = append(k, nums[i])
			ps := permuteUnique(s)
			for _, v := range ps {
				set = append(set, append(k, v...))
			}
		}
	}
	return set
}
