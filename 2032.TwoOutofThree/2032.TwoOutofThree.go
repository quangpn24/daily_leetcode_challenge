package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    res := make([]int, 0)
    m := make(map[int]int)
    for _, n := range nums1 {
        m[n] = 1
    }

    for _, n := range nums2 {
        v, ok := m[n]
        if ok && v == 1 {
            res = append(res, n)
            m[n] = -1
        }
        if !ok {
            m[n] = 2
        }

    }
    for _, n := range nums3 {
        if v, ok := m[n]; ok && (v == 2 || v == 1) {
            res = append(res, n)
            m[n] = -1
        }
    }
    return res
}
func main() {
    nums1 := []int{2, 15, 10, 11, 14, 12, 14, 11, 9, 1}
    nums2 := []int{8, 9, 13, 2, 11, 8}
    nums3 := []int{13, 5, 15, 7, 12, 7, 8, 3, 13, 15}
    fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}
