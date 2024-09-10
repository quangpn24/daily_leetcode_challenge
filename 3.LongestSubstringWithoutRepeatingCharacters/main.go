package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 0

	m := make(map[rune]bool)
	l, r := 0, 0
	for r < len(s) {
		if _, ok := m[rune(s[r])]; !ok {
			m[rune(s[r])] = true
			r++
		} else {
			delete(m, rune(s[l]))
			maxLen = max(maxLen, r-l)
			l++
		}
	}

	return max(maxLen, r-l)
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
}
