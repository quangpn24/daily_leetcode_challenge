package main

import "fmt"

func pivotIndex(nums []int) int {
	sumOfParentArr := 0
	for i := 0; i < len(nums); i++ {
		sumOfParentArr += nums[i]
	}

	sumOfChildArr := 0 // = sum of nums[0] -> nums[ i - 1], i is pivot temp
	for i := 0; i < len(nums); i++ {
		sumOfParentArr -= nums[i]
		if sumOfParentArr == sumOfChildArr {
			return i
		}
		sumOfChildArr += nums[i]
	}
	return -1
}
func main() {
	nums := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(pivotIndex(nums))
}
