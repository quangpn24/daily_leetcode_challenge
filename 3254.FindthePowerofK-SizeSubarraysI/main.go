package main

import "encoding/json"

func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	l, r := 0, 1
	res := make([]int, 0)
	for r < len(nums) {
		for l <= r && l+1 < len(nums) && nums[l] != nums[l+1]-1 {
			l++
		}

		if nums[r] != nums[r-1]+1 {
			l = r - 1
		}
		if r >= k-1 {
			if r-l+1 == k && nums[r] == nums[r-1]+1 {
				res = append(res, nums[r])
			} else {
				res = append(res, -1)
			}
		}
		r++
		if l <= r-k {
			l = r - k + 1
		}
	}
	return res
}
func main() {
	res := resultsArray([]int{3, 2, 3, 2, 3, 2}, 2)
	b, _ := json.Marshal(res)
	println(string(b))
}
