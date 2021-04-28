package main

func searchRange(nums []int, target int) []int {
	//nums := []int{5, 7, 7, 8, 8, 10}
	//              0  1  2  3  4  5

	if len(nums) == 0 {
		return []int{-1, -1}
	}
	//fmt.Println("LEFT")
	i := findLeft(nums, target, 0, len(nums)-1)
	//fmt.Println(i)
	//fmt.Println("RIGHT")
	j := findRight(nums, target, 0, len(nums)-1)
	//fmt.Println(j)

	return []int{i, j}
}

func findLeft(nums []int, target int, i int, j int) int {
	//fmt.Printf("i = %d j = %d\n", i, j)
	if j-i < 2 {
		if nums[i] != target && nums[j] != target {
			return -1
		} else {
			if nums[i] == target {
				return i
			} else {
				return j
			}
		}
	} else {
		m := (i + j) / 2
		if nums[m] >= target {
			return findLeft(nums, target, i, m)
		} else {
			return findLeft(nums, target, m, j)
		}
	}
}

func findRight(nums []int, target int, i int, j int) int {
	//fmt.Printf("i = %d j = %d\n", i, j)
	if j-i < 2 {
		if nums[i] != target && nums[j] != target {
			return -1
		} else {
			if nums[j] == target {
				return j
			} else {
				return i
			}
		}
	} else {
		m := (i + j) / 2
		if nums[m] <= target {
			return findRight(nums, target, m, j)
		} else {
			return findRight(nums, target, i, m)
		}
	}
}
