package main

import (
    "fmt"
)

func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    for (n != 1 && n != -1) && (n%5 == 0 || n%2 == 0 || n%3 == 0) {
        if n%5 == 0 {
            n = n / 5
        }
        if n%2 == 0 {
            n = n / 2
        }
        if n%3 == 0 {
            n = n / 3
        }
    }
    return n == 1
}
func main() {
    fmt.Println(isUgly(-2147483648))
}
