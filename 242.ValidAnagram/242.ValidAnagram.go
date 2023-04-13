package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	origin := make(map[uint8]int) // map contain letters of s
	for i := 0; i < len(s); i++ {
		if v, ok := origin[s[i]]; ok {
			origin[s[i]] = v + 1
		} else {
			origin[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if v, ok := origin[t[i]]; ok {
			origin[t[i]] = v - 1
			if v-1 == 0 {
				delete(origin, t[i])
			}
		} else {
			return false
		}
	}
	return len(origin) == 0
}

func main() {
	s := "cat"
	t := "cat"
	fmt.Println(isAnagram(s, t))
}
