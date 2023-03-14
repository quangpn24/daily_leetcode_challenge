package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapByS := map[uint8]uint8{} //map character in t by s
	mapByT := map[uint8]uint8{} //map character in s by t
	for i := 0; i < len(s); i++ {
		if value, ok := mapByS[s[i]]; ok {
			if value != t[i] {
				return false
			}
		} else {
			if _, ok := mapByT[t[i]]; ok {
				return false
			}
		}
		mapByT[t[i]] = s[i]
		mapByS[s[i]] = t[i]
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
