package main

import (
	"fmt"
	"math"
)

func Max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func rob(nums []int) int {
	n := len(nums)
	result := make([]int, n)
	result[0] = nums[0]
	if n == 1 {
		return nums[0]
	}
	if n >= 2 {
		result[1] = Max(nums[0], nums[1])
	}
	if n >= 3 {
		result[2] = Max(nums[0]+nums[2], nums[1])
	}
	for i := 3; i < n; i++ {
		result[i] = Max(result[i-2], result[i-3]) + nums[i]
	}
	return Max(result[n-1], result[n-2])
}
func main() {
	nums := []int{1, 2, 2, 2}
	fmt.Println(rob(nums))
}
