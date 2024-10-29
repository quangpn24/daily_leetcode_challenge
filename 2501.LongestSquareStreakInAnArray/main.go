package main

import "sort"

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	m := make(map[int]int)

	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num]++
		}
	}

	for _, num := range nums {
		if _, ok := m[num*num]; ok {
			m[num*num] = m[num] + 1
		}
	}

	maxSize := -1
	for _, value := range m {
		maxSize = max(maxSize, value)
	}

	if maxSize == 1 {
		return -1
	}

	return maxSize

}
func main() {
	println(longestSquareStreak([]int{2, 2}))
}
