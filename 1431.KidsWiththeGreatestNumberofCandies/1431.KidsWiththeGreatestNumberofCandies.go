package main

import (
	"fmt"
	"sort"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	sortedCandies := make([]int, len(candies))
	copy(sortedCandies, candies)
	sort.Ints(sortedCandies)
	max := sortedCandies[len(sortedCandies)-1]
	for i, candy := range candies {
		res[i] = candy+extraCandies >= max
	}
	return res
}
func main() {
	candies := []int{12, 1, 12}
	fmt.Println(kidsWithCandies(candies, 10))
}
