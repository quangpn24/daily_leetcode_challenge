package main

import "sort"

func findContentChildren(g []int, s []int) int {
	cnt := 0
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	for _, v := range g {
		for j < len(s) && v > s[j] {
			j++
		}
		if j < len(s) && v <= s[j] {
			j++
			cnt++
		} else if j >= len(s) {
			return cnt
		}
	}
	return cnt
}

func main() {
	println(findContentChildren([]int{1, 2, 3}, []int{3}))
}
