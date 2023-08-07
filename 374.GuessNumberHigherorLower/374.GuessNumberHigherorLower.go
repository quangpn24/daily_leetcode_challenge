package main

import "fmt"

func guess(n int) int {
	pick := 6
	if n == pick {
		return 0
	} else if n < pick {
		return 1
	} else {
		return -1
	}
}
func guessNumber(n int) int {
	l := 0
	r := n
	for {
		mid := (l + r) / 2
		value := guess(mid)
		if value == 0 {
			return mid
		} else if value == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0
}
func main() {
	fmt.Println(guessNumber(6))
}
