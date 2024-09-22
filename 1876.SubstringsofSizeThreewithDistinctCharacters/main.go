package main

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	start, end := 0, 2
	m := make(map[uint8]int)
	m[s[0]]++
	m[s[1]]++
	m[s[2]]++
	ans := 0
	for end < len(s)-1 {
		if len(m) == 3 {
			ans++
		}
		end++
		m[s[end]]++
		m[s[start]]--
		if m[s[start]] == 0 {
			delete(m, s[start])
		}
		start++
	}

	if len(m) == 3 {
		ans++
	}

	return ans
}
func main() {
	println(countGoodSubstrings("xyzzaz"))
}
