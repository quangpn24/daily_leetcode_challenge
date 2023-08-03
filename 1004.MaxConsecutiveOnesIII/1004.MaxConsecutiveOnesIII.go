package main

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	l := 0
	r := 0
	count := 0
	length := 0
	maxLength := -1
	for r < len(nums) {
		if nums[r] == 1 {
			length++
			r++
		} else if count < k {
			length++
			count++
			r++
		} else {
			break
		}
		if length > maxLength {
			maxLength = length
		}
	}
	for r < len(nums) {
		if nums[r] == 0 {
			if nums[l] != 0 {
				length--
				l++
			} else {
				r++
				l++
			}
		} else {
			length++
			r++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(longestOnes(nums, 5))
}
