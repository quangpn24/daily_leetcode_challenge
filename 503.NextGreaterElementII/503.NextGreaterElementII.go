package main

import "fmt"

func nextGreaterElements(nums []int) []int {
    n := len(nums)
    nums = append(nums, nums[0])
    for i := 0; i < n; i++ {
        if nums[i] < nums[i+1] {
            nums[i] = nums[i+1]
        } else {
            nums[i] = -1
        }
    }
    return nums[:n]
}
func main() {
    nums := []int{1, 2, 1}
    fmt.Println(nextGreaterElements(nums))
}
