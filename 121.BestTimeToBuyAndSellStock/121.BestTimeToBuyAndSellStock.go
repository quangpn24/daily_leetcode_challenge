package main

import "math"

// solution 1: sliding window, O(n)
func maxProfit(prices []int) int {
	l, r := 0, 1 // l: buy, r:sell
	profitMax := 0
	for r < len(prices) {
		if prices[r]-prices[l] > profitMax {
			profitMax = prices[r] - prices[l]
		}
		if prices[r] < prices[l] {
			l = r
		}
		r++
	}

	return profitMax
}

// solution 2: greedy, O(n)
func maxProfit2(prices []int) int {
	min := math.MaxInt
	res := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if res < v-min {
			res = v - min
		}
	}
	return res
}
