package main

import (
    "fmt"
)

func nthUglyNumber(n int) int {
    uglys := make([]int, n)
    uglys[0] = 1
    i2, i3, i5 := 0, 0, 0
    for i := 1; i < n; i++ {
        next2, next3, next5 := 2*uglys[i2], 3*uglys[i3], 5*uglys[i5]
        nextUgly := min(next2, next3, next5)
        if nextUgly == next2 {
            i2++
        }
        if nextUgly == next3 {
            i3++
        }
        if nextUgly == next5 {
            i5++
        }
        uglys[i] = nextUgly
    }
    return uglys[n-1]
}
func min(a int, b int, c int) int {
    tmp := a
    if a > b {
        tmp = b
    }
    if tmp > c {
        return c
    }
    return tmp
}
func main() {
    fmt.Println(nthUglyNumber(1352))
}
