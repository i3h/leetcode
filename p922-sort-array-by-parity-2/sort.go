package main

func sortArrayByParityII(nums []int) []int {
	odds := []int{}
	evens := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evens = append(evens, nums[i])
		} else {
			odds = append(odds, nums[i])
		}
	}
	for i := 0; i < len(nums); i += 2 {
		nums[i] = evens[i/2]
	}
	for i := 1; i < len(nums); i += 2 {
		nums[i] = odds[i/2]
	}

	return nums
}
