package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		return n == 0 || flowerbed[0] == 0
	}
	next := 1
	pre := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[pre] == 0 && flowerbed[i] == 0 && flowerbed[next] == 0 {
			flowerbed[i] = 1
			n--
		}
		if i != len(flowerbed)-2 {
			next++
		}
		pre = i
	}
	return n <= 0
}
func main() {
	flowerbed := []int{1}
	n := 0
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
