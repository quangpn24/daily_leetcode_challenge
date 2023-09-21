package main

import "fmt"

func plusOne(digits []int) []int {
    cnt := 0
    n := len(digits)
    if digits[n-1]+1 >= 10 {
        digits[n-1] = (digits[n-1] + 1) % 10
        cnt = 1
    } else {
        digits[n-1] = digits[n-1] + 1
        return digits
    }
    for i := n - 2; i >= 0; i-- {
        if digits[i]+cnt >= 10 {
            digits[i] = (digits[i] + cnt) % 10
            cnt = 1
        } else {
            digits[i] = digits[i] + cnt
            return digits
        }
    }
    if cnt == 1 {
        return append([]int{1}, digits...)
    }
    return digits
}
func main() {
    fmt.Println(plusOne([]int{9, 9}))
}
