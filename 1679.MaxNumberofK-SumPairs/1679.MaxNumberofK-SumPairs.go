package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	res := 0
	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] > k {
			r--
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			res++
			l++
			r--
		}
	}
	return res
}
func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}
