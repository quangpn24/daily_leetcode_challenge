package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	return int64(count(nums, lower) - count(nums, upper+1))
}

func count(nums []int, lower int) int {
	cnt := 0
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]+nums[r] < lower {
			l++
		}
		println(r, l)
		cnt += r - l
		r--
	}
	return cnt
}

func main() {
	println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
}
