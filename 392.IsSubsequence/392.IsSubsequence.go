package main

// use two pointer
func isSubsequence(s string, t string) bool {
	indexT := 0
	indexS := 0
	for indexS < len(s) && indexT < len(t) {
		if s[indexS] == t[indexT] {
			indexS++
			indexT++
		} else {
			indexT++
		}
	}
	return indexS == len(s)
}
