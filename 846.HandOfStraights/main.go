package main

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	m := make(map[int]int)
	for _, v := range hand {
		m[v]++
	}

	sort.Ints(hand)
	for _, v := range hand {
		if m[v] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			if m[v+i] == 0 {
				return false
			}
			m[v+i]--
		}
	}
	return true
}

func main() {
	println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}
