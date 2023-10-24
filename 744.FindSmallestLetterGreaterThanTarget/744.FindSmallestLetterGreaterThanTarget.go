package main

import (
	"fmt"
	"math"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	min := math.MaxInt
	res := letters[0]
	for _, letter := range letters {
		if letter > target && int(letter)-int(target) < min {
			min = int(letter) - int(target)
			res = letter
		}
	}
	return res
}
func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
}
