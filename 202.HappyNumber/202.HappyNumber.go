package main

import (
    "fmt"
    "strconv"
)

func isHappy(n int) bool {
    m := make(map[int]bool)
    for n != 1 {
        str := strconv.Itoa(n)
        sum := 0
        for _, s := range str {
            tmp, _ := strconv.Atoi(string(s))
            sum += tmp * tmp
        }
        if sum == 1 {
            return true
        }
        if _, ok := m[sum]; ok {
            return false
        } else {
            m[sum] = true
        }
        n = sum
    }
    return true
}
func main() {
    fmt.Println(isHappy(19))
}
