package main

import (
	"fmt"
	"math"
)

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	minLen := int(math.Min(float64(len(word1)), float64(len(word2))))
	for i := 0; i < minLen; i++ {
		if i < len(word1) {
			res += string(word1[i])
		}
		if i < len(word2) {
			res += string(word2[i])
		}
	}
	if len(word1) > len(word2) {
		res += string(word1[len(word2):len(word1)])
	}
	if len(word1) < len(word2) {
		res += string(word2[len(word1):len(word2)])
	}
	return res
}
func main() {
	word1 := "ab"
	word2 := "pqr"
	fmt.Println(mergeAlternately(word1, word2))
}
