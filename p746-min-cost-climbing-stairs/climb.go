package main

var cache = map[int]int{0: 0}
var first_time bool = true

func minCostClimbingStairs(cost []int) int {
	if first_time {
		cost = append([]int{0}, cost...)
		first_time = false
	}
	l := len(cost)
	if l == 0 {
		return 0
	}
	if l == 1 {
		cache[1] = cost[0]
		return cost[0]
	}
	if _, ok := cache[l]; ok {
		return cache[l]
	}
	cache[l] = min(minCostClimbingStairs(cost[1:])+cost[0], minCostClimbingStairs(cost[2:])+cost[0])
	return minCostClimbingStairs(cost)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
