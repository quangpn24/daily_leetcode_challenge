package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)

	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	testCases := []int{121, -121, 0, 123456654321, 11111123213, 111111}
	for _, tc := range testCases {
		check := isPalindrome(tc)
		if check {
			fmt.Printf("%v is a palidrome number\n", tc)
		} else {
			fmt.Printf("%v is not a palidrome number\n", tc)
		}
	}
}
