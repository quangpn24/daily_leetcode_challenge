package main

import "fmt"

func missingNumber(nums []int) int {
    n := len(nums)
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return n*(n+1)/2 - sum
}
func main() {
    fmt.Println(missingNumber([]int{3, 1, 0, 2}))
}
