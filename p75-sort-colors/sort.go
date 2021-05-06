package main

func sortColors(nums []int) {
	counts := []int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}
	j := 0
	for i := 0; i < len(counts); i++ {
		for counts[i] > 0 {
			nums[j] = i
			j++
			counts[i]--
		}
	}
}
