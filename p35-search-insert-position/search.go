package main

func searchInsert(nums []int, target int) int {
	//fmt.Println("target = ", target)
	l := len(nums)
	i := 0
	j := l - 1
	m := (i + j) / 2

	for i < j-1 {
		//fmt.Printf("i = %d j = %d m = %d\n", nums[i], nums[j], nums[m])
		if target < nums[m] {
			j = m
			m = (i + j) / 2
		} else {
			i = m
			m = (i + j) / 2
		}
	}
	//fmt.Printf("i = %d j = %d m = %d\n", nums[i], nums[j], nums[m])

	if i == j {
		if target <= nums[i] {
			return i
		} else {
			return i + 1
		}
	} else {
		if target <= nums[i] {
			return i
		} else if target <= nums[j] {
			return j
		} else {
			return j + 1
		}
	}
}
