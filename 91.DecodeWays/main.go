package main

func numDecodings(s string) int {
	m := make([]int, len(s))
	m[0] = 1
	if s[0] == '0' {
		return 0
	}
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			m[i] = m[i-1]
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			if i == 1 {
				m[i]++
			} else {
				m[i] += m[i-2]
			}
		}
	}
	return m[len(s)-1]
}
func main() {
	println(numDecodings("1123"))
}
