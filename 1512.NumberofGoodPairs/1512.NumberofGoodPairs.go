package main

import "fmt"

func numIdenticalPairs(nums []int) int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }

    res := 0
    for _, value := range m {
        if value > 1 {
            res += value * (value - 1) / 2
        }
    }
    return res
}

func main() {
    fmt.Println(numIdenticalPairs([]int{1}))
}
