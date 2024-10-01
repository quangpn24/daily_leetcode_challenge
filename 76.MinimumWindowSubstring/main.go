package main

func minWindow(s string, t string) string {
	m := make(map[uint8]int)
	for _, c := range t {
		m[uint8(c)]++
	}

	start, end := 0, 0
	minWd := len(s) + 1
	res := ""
	count := 0
	for end < len(s) {
		//char exist in target
		if m[s[end]] > 0 {
			count++
		}
		// if char not exist, value is negative number
		m[s[end]]--

		for count == len(t) {
			if minWd > end-start+1 {
				minWd = end - start + 1
				res = s[start : end+1]
			}
			m[s[start]]++
			if m[s[start]] > 0 {
				count--
			}
			start++
		}
		end++
	}

	return res
}

func main() {
	println(minWindow("ADOBECODEBANC", "ABC"))
}
