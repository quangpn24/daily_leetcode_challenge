package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'o' || s[i] == 'i' || s[i] == 'u' {
			count++
		}
	}
	l := 0
	r := k
	max := count
	for r < len(s) {
		if count == k {
			return count
		}
		if s[r] == 'a' || s[r] == 'e' || s[r] == 'o' || s[r] == 'i' || s[r] == 'u' {
			count++
		}
		if l != r && (s[l] == 'a' || s[l] == 'e' || s[l] == 'o' || s[l] == 'i' || s[l] == 'u') {
			count--
		}
		if count > max {
			max = count
		}
		r++
		l++
	}
	return max
}
func main() {
	s := ""
	k := 3
	fmt.Println(maxVowels(s, k))
}
