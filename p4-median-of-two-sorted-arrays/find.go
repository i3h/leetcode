package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	i := l/2 - (1 - l%2)
	j := l / 2

	if l1 == 0 {
		return float64(nums2[i]+nums2[j]) / 2
	} else if l2 == 0 {
		return float64(nums1[i]+nums1[j]) / 2
	} else {
		var m, m0 float64
		var p1, p2 int
		for {
			if p1 >= l1 {
				p2++
				m0 = m
				m = float64(nums2[p2-1])
				if p1+p2 > j {
					break
				}
			} else if p2 >= l2 {
				p1++
				m0 = m
				m = float64(nums1[p1-1])
				if p1+p2 > j {
					break
				}
			} else {
				if nums1[p1] < nums2[p2] {
					p1++
					m0 = m
					m = float64(nums1[p1-1])
					if p1+p2 > j {
						break
					}
				} else {
					p2++
					m0 = m
					m = float64(nums2[p2-1])
					if p1+p2 > j {
						break
					}
				}
			}
		}
		if i == j {
			return m
		} else {
			return (m + m0) / 2
		}
	}
}
