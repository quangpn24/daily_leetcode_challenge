package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAverage := float64(sum) / float64(k)
	l := 0
	r := k
	for r < len(nums) {
		sum = sum + nums[r] - nums[l]
		if float64(sum)/float64(k) > maxAverage {
			maxAverage = float64(sum) / float64(k)
		}
		r++
		l++
	}
	return maxAverage
}
func main() {
	nums := []int{5}
	k := 1
	fmt.Println(findMaxAverage(nums, k))
}
