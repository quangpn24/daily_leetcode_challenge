package main

import (
	"fmt"
	"strconv"
)

func addLetter(origin []string, letters []string) []string {
	res := make([]string, 0)
	for _, s := range origin {
		for _, letter := range letters {
			res = append(res, s+letter)
		}
	}
	return res
}
func letterCombinations(digits string) []string {
	mDigits := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}

	res := []string{}
	for index, digit := range digits {
		num, _ := strconv.Atoi(string(digit))
		if index == 0 {
			res = append(res, mDigits[num]...)
		} else {
			res = addLetter(res, mDigits[num])
		}
	}
	return res
}
func main() {
	digits := ""
	fmt.Println(letterCombinations(digits))
}
