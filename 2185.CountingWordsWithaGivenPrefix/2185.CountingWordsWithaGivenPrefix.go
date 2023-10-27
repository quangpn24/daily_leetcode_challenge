package main

import "fmt"

func prefixCount(words []string, pref string) int {
	cnt := 0
	for _, w := range words {
		if len(w) <= len(pref) && w[:len(pref)] == pref {
			cnt++
		}
	}
	return cnt
}
func main() {
	fmt.Println(prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
}
