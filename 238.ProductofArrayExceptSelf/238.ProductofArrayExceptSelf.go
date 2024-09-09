package main

import "fmt"

func productExceptSelf(nums []int) []int {
	product := 1
	prefix := make([]int, len(nums)+1)
	prefix[0] = 1
	res := make([]int, len(nums))
	for i, num := range nums {
		product *= num
		prefix[i+1] = product
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = prefix[i] * suffix
		suffix *= nums[i]
	}

	return res
}
func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
