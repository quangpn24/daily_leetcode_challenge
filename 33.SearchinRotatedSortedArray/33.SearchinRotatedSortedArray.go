package main

import "fmt"

func search(nums []int, target int) int {
    pivot := len(nums) - 1
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            pivot = i
            break
        }
    }
    index := -1
    if target > nums[pivot] {
        index = BinarySearch(nums[:pivot], target)
    } else if target < nums[pivot] {
        index = BinarySearch(nums[pivot+1:], target)
        if index != -1 {
            index = index + pivot + 1
        }
    } else {
        return pivot
    }
    return index
}
func BinarySearch(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        m := (nums[l] + nums[r]) / 2
        if nums[m] == target {
            return m
        } else if nums[m] < target {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return -1
}
func main() {
    fmt.Println(search([]int{1, 3}, 1))
}
