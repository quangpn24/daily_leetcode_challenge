package main

import (
    "fmt"
    "sort"
)

func minMoves2(nums []int) int {
    sort.Ints(nums)

    res := 0
    mid := nums[len(nums)/2]
    for _, num := range nums {
        if num < mid {
            res += mid - num
        } else {
            res += num - mid
        }
    }
    return res
}
func main() {
    nums := []int{1, 10, 2, 9}
    fmt.Println(minMoves2(nums))
}
