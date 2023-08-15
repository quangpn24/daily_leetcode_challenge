package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := []int{}
	for _, spell := range spells {
		l, r := 0, len(potions)-1
		count := len(potions)
		for l <= r {
			mid := (l + r) / 2
			if int64(potions[mid]*spell) >= success {
				r = mid - 1
			} else {
				count = count - (mid - l + 1)
				l = mid + 1
			}
		}
		res = append(res, count)
	}
	return res
}
func main() {
	spells := []int{3, 1, 2}
	potions := []int{8, 5, 8}
	success := 16
	fmt.Println(successfulPairs(spells, potions, int64(success)))
}
