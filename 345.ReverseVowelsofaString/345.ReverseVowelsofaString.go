package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1
	arr := strings.Split(s, "")
	for left < right {
		for left < len(arr) && strings.ToLower(arr[left]) != "a" && strings.ToLower(arr[left]) != "e" && strings.ToLower(arr[left]) != "o" && strings.ToLower(arr[left]) != "i" && strings.ToLower(arr[left]) != "u" {
			left++
		}
		for right >= 0 && strings.ToLower(arr[right]) != "a" && strings.ToLower(arr[right]) != "e" && strings.ToLower(arr[right]) != "o" && strings.ToLower(arr[right]) != "i" && strings.ToLower(arr[right]) != "u" {
			right--
		}

		//swap
		if left < right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			left++
			right--
		}
	}
	return strings.Join(arr, "")
}
func main() {
	s := "aA"
	fmt.Println(reverseVowels(s))
}
