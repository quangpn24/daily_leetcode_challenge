package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

func removeStars(s string) string {
	myStack := utils.Stack[byte]{}
	res := ""
	for _, b := range s {
		myStack.Push(byte(b))
	}
	count := 0
	for !myStack.IsEmpty() {
		ch := myStack.Pop()
		if ch == '*' {
			count++
		} else if count == 0 {
			res = string(ch) + res
		} else {
			count--
		}
	}
	return res
}
func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}
