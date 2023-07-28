package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := []uint8{}
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	target := strs[0]
	index := 0
	for index < len(target) {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || target[index] != strs[i][index] {
				return string(res)
			}
		}
		res = append(res, target[index])
		index++
	}
	return string(res)
}
func main() {
	strs := []string{"abc"}
	fmt.Println(longestCommonPrefix(strs))
}
