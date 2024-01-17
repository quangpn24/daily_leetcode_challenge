package main

import "fmt"

func canBeTypedWords(text string, brokenLetters string) int {
	letters := make([]bool, 26)
	for _, l := range brokenLetters {
		letters[l-'a'] = true
	}
	cnt := 1
	flag := false
	for _, l := range text {
		if l == ' ' {
			flag = false
			cnt++
			continue
		}
		if !flag && letters[l-'a'] {
			cnt--
			flag = true
		}
	}
	return cnt
}
func main() {
	fmt.Println(canBeTypedWords("leet code", "l"))
}
