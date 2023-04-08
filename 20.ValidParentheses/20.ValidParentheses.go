package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

func findChar(arr []rune, c rune) int {
	for i, r := range arr {
		if c == r {
			return i
		}
	}
	return -1
}

func isValid(s string) bool {
	openChar := []rune{'(', '{', '['}
	closeChar := []rune{')', '}', ']'}
	stack := utils.Stack[rune]{}
	for _, c := range s {
		index := findChar(openChar, c)
		if index != -1 {
			stack.Push(c)
		}

		indexClose := findChar(closeChar, c)
		if indexClose != -1 {
			if indexClose == findChar(openChar, stack.Top()) {
				stack.Pop()
			} else {
				return false
			}

		}
	}
	return stack.IsEmpty()
}

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}
