package main

func permute(nums []int) [][]int {
	var set [][]int
	if len(nums) == 1 {
		return [][]int{nums}
	} else {
		for i := 0; i < len(nums); i++ {
			s := []int{}
			k := []int{}
			s = append(s, nums[:i]...)
			s = append(s, nums[i+1:]...)
			k = append(k, nums[i])
			ps := permute(s)
			for _, v := range ps {
				set = append(set, append(k, v...))
			}
		}
	}
	return set
}
