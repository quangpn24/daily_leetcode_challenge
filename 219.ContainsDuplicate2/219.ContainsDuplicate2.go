package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	mapNums := map[int]int{}
	for i, num := range nums {
		if value, ok := mapNums[num]; ok && int(math.Abs(float64(value-i))) <= k {
			return true
		} else {
			mapNums[num] = i
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
}
