package main

import (
    "fmt"
)

func firstUniqChar(s string) int {
    myMap := make(map[int32]int)
    for _, c := range s {
        myMap[c]++
    }
    for i, c := range s {
        if myMap[c] == 1 {
            return i
        }
    }
    return -1
}
func main() {
    fmt.Println(firstUniqChar("leetcode"))
}
