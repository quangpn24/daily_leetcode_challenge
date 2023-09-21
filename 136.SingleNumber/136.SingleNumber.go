package main

import "fmt"

func singleNumber(nums []int) int {
    mMap := make(map[int]int)
    for _, num := range nums {
        if v, ok := mMap[num]; ok && v == 2 {
            delete(mMap, num)
        } else {
            mMap[num]++
        }
    }
    for k, v := range mMap {
        if v == 1 {
            return k
        }
    }
    return -1
}
func main() {
    fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
