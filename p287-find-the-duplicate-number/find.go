package main

func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		} else {
			m[nums[i]] = 1
		}
	}
	return 0
}
