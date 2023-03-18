package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		if height[l] < height[r] {
			//If area < max, set max
			if height[l]*(r-l) > res {
				res = height[l] * (r - l)
			}
			l++
		} else {
			//If area > max, set max
			if height[r]*(r-l) > res {
				res = height[r] * (r - l)
			}
			r--
		}
	}
	return res
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
