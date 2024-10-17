package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtrack(&res, []int{}, candidates, target, 0, 0)
	return res
}

func backtrack(res *[][]int, combination []int, candidates []int, target int, sum int, i int) {
	if sum == target {
		tmp := make([]int, len(combination))
		copy(tmp, combination)
		*res = append(*res, tmp)
	} else if sum > target {
		return
	}

	for j := i; j < len(candidates); j++ {
		if j > i && candidates[j] == candidates[j-1] {
			continue
		}
		combination = append(combination, candidates[j])
		sum += candidates[j]
		backtrack(res, combination, candidates, target, sum, j+1)
		combination = combination[:len(combination)-1]
		sum -= candidates[j]
	}
}

func main() {
	res := combinationSum2([]int{1, 1, 6, 7}, 8)
	for _, re := range res {
		for _, r := range re {
			fmt.Print(r, " ")
		}
		fmt.Println()
	}
}
