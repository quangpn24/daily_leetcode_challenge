package main

import (
	"fmt"
	"strings"
)

// use two pointer
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		flag := true
		if !(s[l] >= 97 && s[l] <= 122) && !(s[l] >= 48 && s[l] <= 57) {
			l++
			flag = false
		}
		if !(s[r] >= 97 && s[r] <= 122) && !(s[r] >= 48 && s[r] <= 57) {
			r--
			flag = false
		}
		if !flag {
			continue
		}
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
