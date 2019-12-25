package main

// Hash table
func twoSum(nums []int, target int) []int {
	indexs := make([]int, 2)

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		com := target - nums[i]
		if val, ok := m[com]; ok && i != val {
			indexs[0] = i
			indexs[1] = val
			break
		}
		m[nums[i]] = i
	}

	return indexs
}
