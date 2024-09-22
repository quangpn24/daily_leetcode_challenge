package main

func findSubstring(s string, words []string) []int {
	n := len(words[0]) * len(words)
	ans := make([]int, 0)

	mWord := make(map[string]int)
	wordLen := len(words[0])
	for _, word := range words {
		mWord[word]++
	}

	for i := 0; i < wordLen; i++ {
		for j := i; j <= len(s)-wordLen; j += wordLen {
			if j+n > len(s) {
				break
			}
			m := makeMap(s[j:j+n], len(words[0]))
			if check(m, mWord) {
				ans = append(ans, j)
			}
		}
	}
	return ans
}

func makeMap(s string, k int) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(s); i += k {
		m[s[i:i+k]]++
	}
	return m
}

func check(m map[string]int, mWords map[string]int) bool {
	// compare two maps
	if len(m) != len(mWords) {
		return false
	}
	for k, v := range m {
		if mWords[k] != v {
			return false
		}
	}

	return true
}

func main() {
	//ans := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	//for _, v := range ans {
	//	print(v, " ")
	//}

	//ans := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	//for _, v := range ans {
	//	print(v, " ")
	//}

	//ans := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	//for _, v := range ans {
	//	print(v, " ")
	//}

	ans := findSubstring("a", []string{"a"})
	for _, v := range ans {
		print(v, " ")
	}
}
