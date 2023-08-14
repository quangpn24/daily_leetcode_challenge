package main

import "fmt"

func findDuplicate(nums []int) int {
	myMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := myMap[num]; ok {
			return num
		} else {
			myMap[num] = true
		}
	}
	return 0
}
func main() {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println(findDuplicate(nums))
}
