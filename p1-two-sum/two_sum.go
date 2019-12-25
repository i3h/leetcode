package main

func twoSum(nums []int, target int) []int {
	indexs := make([]int, 2)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indexs[0] = i
				indexs[1] = j
				break
			}
		}
	}

	return indexs
}
