package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]int, len(nums))
	sort.Ints(nums)
	backtrack(&res, []int{}, nums, used)
	return res
}

func backtrack(res *[][]int, permutations []int, nums []int, used []int) {
	if len(permutations) == len(nums) {
		tmp := make([]int, len(permutations))
		copy(tmp, permutations)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == 1 || (i > 0 && nums[i] == nums[i-1] && used[i-1] == 0) {
			continue
		}
		permutations = append(permutations, nums[i])
		used[i] = 1
		backtrack(res, permutations, nums, used)
		used[i] = 0
		permutations = permutations[:len(permutations)-1]
	}
}
func main() {
	res := permuteUnique([]int{1, 1, 2, 2})
	for _, re := range res {
		for _, r := range re {
			print(r, " ")
		}
		println("")
	}
}
