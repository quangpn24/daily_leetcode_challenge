package main

import "sort"

func distinctAverages(nums []int) int {
    sort.Ints(nums)
    mMap := make(map[float64]bool)
    l, r := 0, len(nums)-1
    for l < r {
        average := float64(nums[l]+nums[r]) / 2
        mMap[average] = true
        l++
        r--
    }
    return len(mMap)
}
func main() {
    println(distinctAverages([]int{4, 1, 4, 0, 3, 5}))
}
