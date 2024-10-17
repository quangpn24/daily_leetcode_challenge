package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]int, len(nums))
	backtrack(nums, &res, []int{}, used)
	return res
}

func backtrack(nums []int, res *[][]int, permutations []int, used []int) {
	if len(permutations) == len(nums) {
		tmp := make([]int, len(permutations))
		copy(tmp, permutations)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == 0 {
			permutations = append(permutations, nums[i])
			used[i] = 1
			backtrack(nums, res, permutations, used)
			used[i] = 0
			permutations = permutations[:len(permutations)-1]
		}
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
