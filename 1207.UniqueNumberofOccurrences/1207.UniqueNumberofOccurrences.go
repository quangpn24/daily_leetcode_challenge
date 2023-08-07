package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	myMap := make(map[int]int)
	for _, e := range arr {
		if v, ok := myMap[e]; ok {
			myMap[e] = v + 1
		} else {
			myMap[e] = 1
		}
	}

	mapTmp := make(map[int]bool)
	for _, value := range myMap {
		if _, ok := mapTmp[value]; ok {
			return false
		} else {
			mapTmp[value] = true
		}

	}
	return true
}
func main() {
	arr := []int{1, 2, 3, 1, 1, 2}
	fmt.Println(uniqueOccurrences(arr))
}
