package main

import (
	"math"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	for i := 1; i < len(items); i++ {
		items[i][1] = int(math.Max(float64(items[i-1][1]), float64(items[i][1])))
	}

	res := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		l, r := 0, len(items)-1
		for l < r {
			mid := l + (r-l+1)/2

			if items[mid][0] > queries[i] {
				r = mid - 1
			} else {
				l = mid
			}
		}
		if items[l][0] <= queries[i] {
			res[i] = items[l][1]
		}
	}
	return res
}

func main() {
	res := maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6})
	for _, re := range res {
		println(re)
	}
}
