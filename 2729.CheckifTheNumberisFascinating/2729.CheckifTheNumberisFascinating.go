package main

import (
	"fmt"
	"strconv"
)

func isFascinating(n int) bool {
	arr := make([]bool, 10)
	strNum := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
	if len(strNum) != 9 {
		return false
	}
	for _, s := range strNum {
		index, _ := strconv.Atoi(string(s))
		if arr[index] || index == 0 {
			return false
		} else {
			arr[index] = true
		}
	}
	return true
}
func main() {
	fmt.Println(isFascinating(267))
}
