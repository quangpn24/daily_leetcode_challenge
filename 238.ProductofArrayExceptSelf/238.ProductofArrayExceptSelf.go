package main

import "fmt"

func productExceptSelf(nums []int) []int {
	product := 1
	numOfZero := 0
	for _, num := range nums {
		if num == 0 {
			numOfZero++
			if numOfZero > 1 {
				return make([]int, len(nums))
			}
		} else {
			product = product * num
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = product
		} else {
			nums[i] = product * (1 - numOfZero) / nums[i]
		}
	}
	return nums
}
func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
