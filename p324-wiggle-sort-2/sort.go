package main

func wiggleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		//fmt.Println("i = ", i)
		//fmt.Println("before: ", nums)
		if i%2 == 0 {
			//fmt.Println("search greater")
			for j := i + 1; j < l; j++ {
				if nums[j] > nums[i] {
					//fmt.Println("num i: ", nums[i])
					//fmt.Println("num j: ", nums[j])
					swap(nums, i+1, j)
					break
				}
			}
		} else {
			//fmt.Println("search less")
			for j := i + 1; j < l; j++ {
				if nums[j] < nums[i] {
					//fmt.Println("num i: ", nums[i])
					//fmt.Println("num j: ", nums[j])
					swap(nums, i+1, j)
					break
				}
			}
		}
		//fmt.Println("after: ", nums)
	}
}

func swap(nums []int, i int, j int) {
	t := nums[j]
	nums[j] = nums[i]
	nums[i] = t
}
