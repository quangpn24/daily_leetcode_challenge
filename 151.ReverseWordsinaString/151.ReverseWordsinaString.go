package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := ""
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			res += arr[i] + " "
		}
	}
	return res[:len(res)-1]
}
func main() {
	s := "hello     world"
	fmt.Println(reverseWords(s))
}
