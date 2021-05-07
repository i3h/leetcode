package main

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1
	for i < j {
		if A[i]%2 == 0 {
			i++
		} else {
			swap(A, i, j)
			j--
		}
	}

	return A
}

func swap(nums []int, i int, j int) {
	t := nums[j]
	nums[j] = nums[i]
	nums[i] = t
}
