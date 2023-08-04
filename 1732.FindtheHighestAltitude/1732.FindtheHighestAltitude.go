package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	highestAltitude := 0
	currentAltitude := 0
	for _, g := range gain {
		currentAltitude += g
		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}
	return highestAltitude
}
func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}
