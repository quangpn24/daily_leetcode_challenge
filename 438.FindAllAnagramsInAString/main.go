package main

func findAnagrams(s string, p string) []int {
	mTarget := make(map[rune]int)
	for _, c := range p {
		mTarget[c]++
	}

	n := len(p)

	start, end := 0, 0
	res := make([]int, 0)
	m := make(map[rune]int)

	for end < len(s) {
		m[rune(s[end])]++
		if end-start+1 == n {
			//check 2 map
			if compareTwoMap(m, mTarget) {
				res = append(res, start)
			}
			m[rune(s[start])]--
			if m[rune(s[start])] == 0 {
				delete(m, rune(s[start]))
			}
			start++
		}
		end++
	}

	return res
}

func compareTwoMap(m1 map[rune]int, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok && v1 == v2 {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	res := findAnagrams("abab", "ab")
	for _, re := range res {
		print(re, " ")
	}
}
