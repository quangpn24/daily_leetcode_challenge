package main

import "fmt"

func runningSum(nums []int) []int {
	sum := 0 //  = sum of nums[0] -> nums[i]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
