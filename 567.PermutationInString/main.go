package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1arr := [26]int{}
	s2arr := [26]int{}
	for i := 0; i < len(s1); i++ {
		s1arr[s1[i]-'a']++
	}

	start, end := 0, 0
	for end < len(s2) {
		s2arr[s2[end]-'a']++
		if end-start+1 == len(s1) {
			if s1arr == s2arr {
				return true
			}
			s2arr[s2[start]-'a']--
			start++
		}
		end++
	}

	return s1arr == s2arr
}

func main() {
	println(checkInclusion("abc", "eidbaooo"))
}
