package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	var target byte
	num := 0
	index := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] != target {
			if num > 1 {
				for _, b := range strconv.Itoa(num) {
					chars[index] = byte(b)
					index++
				}
			}
			target = chars[i]
			num = 1
			chars[index] = target
			index++
		} else {
			num++
		}
	}

	if num > 1 {
		for _, b := range strconv.Itoa(num) {
			chars[index] = byte(b)
			index++
		}
	}
	chars = chars[:index]
	return len(chars)
}
func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
	fmt.Println(chars)
}
