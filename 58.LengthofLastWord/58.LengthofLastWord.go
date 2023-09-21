package main

import (
    "fmt"
    "strings"
)

func lengthOfLastWord(s string) int {
    arr := strings.Split(s, " ")
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] != "" {
            return len(arr[i])
        }
    }
    return 0
}
func main() {
    fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
