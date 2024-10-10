package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	cnt := 0
	sort.Ints(people)

	l, r := 0, len(people)-1

	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		cnt++
		r--
	}

	return cnt
}

func main() {
	println(numRescueBoats([]int{5, 1, 4, 2}, 6))
}
