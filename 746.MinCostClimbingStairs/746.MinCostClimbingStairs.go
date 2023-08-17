package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	res := make([]int, len(cost)+1)
	res[0] = cost[0]
	res[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		if res[i-1] > res[i-2] {
			res[i] = res[i-2] + cost[i]
		} else {
			res[i] = res[i-1] + cost[i]
		}
	}
	if res[len(cost)-1] > res[len(cost)-2] {
		return res[len(cost)-2]
	} else {
		return res[len(cost)-1]
	}
}
func main() {
	cost := []int{1, 100}
	fmt.Println(minCostClimbingStairs(cost))
}
