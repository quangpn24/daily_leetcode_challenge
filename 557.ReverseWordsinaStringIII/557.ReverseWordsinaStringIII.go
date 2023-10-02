package main

import (
    "fmt"
)

func reverseWords(s string) string {
    arrBytes := []byte(s)
    index := 0
    for i := 0; i <= len(arrBytes); i++ {
        if i == len(arrBytes) || arrBytes[i] == ' ' {
            reverse(arrBytes, index, i-1)
            index = i + 1
        }
    }
    return string(arrBytes)
}
func reverse(s []byte, l int, r int) {
    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
}
func main() {
    fmt.Println(reverseWords("Let's take LeetCode contest"))
}
