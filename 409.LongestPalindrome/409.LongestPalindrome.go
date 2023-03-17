package main

func longestPalindrome(s string) int {
	letterMap := map[rune]int{}

	for _, c := range s {
		if _, ok := letterMap[c]; ok {
			letterMap[c]++
		} else {
			letterMap[c] = 1
		}
	}
	res := 0
	existOdd := false
	for _, v := range letterMap {
		if v%2 == 0 {
			res += v
		} else {
			existOdd = true
			res += v - 1
		}
	}
	if existOdd {
		res++
	}
	return res
}
