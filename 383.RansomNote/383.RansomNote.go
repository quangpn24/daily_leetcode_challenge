package main

func canConstruct(ransomNote string, magazine string) bool {
	myMap := map[rune]int{}
	// insert letters of magazine into map
	for _, c := range magazine {
		if v, ok := myMap[c]; ok {
			myMap[c] = v + 1
		} else {
			myMap[c] = 1
		}
	}

	for _, c := range ransomNote {
		if v, ok := myMap[c]; ok {
			myMap[c] = v - 1
			if v == 1 {
				delete(myMap, c)
			}
		} else {
			return false
		}
	}
	return true
}
