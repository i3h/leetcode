package main

func sortArray(nums []int) []int {
	l := len(nums)
	if l == 0 || l == 1 {
		return nums
	}
	if l == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return nums
	}

	lo := 0
	hi := l - 2
	pi := l - 1

	//fmt.Printf("pi = %d lo = %d hi = %d\n", pi, lo, hi)

	for lo < hi {
		if nums[lo] < nums[pi] {
			lo++
		}
		if nums[hi] > nums[pi] {
			hi--
		}
		if lo < hi && nums[lo] >= nums[pi] && nums[hi] <= nums[pi] {
			/*
				fmt.Println("swap begin")
				fmt.Printf("lo = %d hi = %d pi = %d\n", lo, hi, pi)
				fmt.Println(nums)
			*/
			nums[lo], nums[hi] = nums[hi], nums[lo]
			/*
				fmt.Println(nums)
				fmt.Printf("lo = %d hi = %d pi = %d\n", lo, hi, pi)
				fmt.Println("swap end")
			*/
		}
	}

	if nums[lo] < nums[pi] {
		nums[lo+1], nums[pi] = nums[pi], nums[lo+1]
		pi = lo + 1
	} else {
		nums[lo], nums[pi] = nums[pi], nums[lo]
		pi = lo
	}
	//nums[lo+1], nums[pi] = nums[pi], nums[lo+1]

	//fmt.Println("left:", nums[:pi])
	//fmt.Println("righ:", nums[pi+1:])
	sortArray(nums[:pi])
	sortArray(nums[pi+1:])

	return nums
}
