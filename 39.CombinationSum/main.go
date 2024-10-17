package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	backtrack(&res, []int{}, candidates, target, 0)
	return res
}

func backtrack(res *[][]int, combination []int, candidates []int, target int, sum int) {
	if sum == target {
		tmp := make([]int, len(combination))
		copy(tmp, combination)
		*res = append(*res, tmp)
	} else if sum > target {
		return
	}

	for _, candidate := range candidates {
		if len(combination) == 0 || combination[len(combination)-1] <= candidate {
			combination = append(combination, candidate)
			sum += candidate
			backtrack(res, combination, candidates, target, sum)
			combination = combination[:len(combination)-1]
			sum -= candidate
		}
	}
}

func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	for _, re := range res {
		for _, r := range re {
			fmt.Print(r, " ")
		}
		fmt.Println()
	}
}
