package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	var i, j, k int
	for k < m+n {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				nums3[k] = nums1[i]
				k++
				i++
			} else {
				nums3[k] = nums2[j]
				k++
				j++
			}
		}
		if i == m {
			for j < n {
				nums3[k] = nums2[j]
				k++
				j++
			}
		}
		if j == n {
			for i < m {
				nums3[k] = nums1[i]
				k++
				i++
			}
		}
	}
	//fmt.Println("nums3:", nums3)
	copy(nums1, nums3)
}
