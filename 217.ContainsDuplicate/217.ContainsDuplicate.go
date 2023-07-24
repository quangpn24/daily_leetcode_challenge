package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mapNums := map[int]bool{}
	for _, num := range nums {
		if _, ok := mapNums[num]; ok {
			return true
		} else {
			mapNums[num] = true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
