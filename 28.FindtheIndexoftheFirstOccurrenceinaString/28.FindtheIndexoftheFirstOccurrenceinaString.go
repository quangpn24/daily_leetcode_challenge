package main

import "fmt"

func strStr(haystack string, needle string) int {
    n := len(haystack)
    for i := 0; i < n; i++ {
        if haystack[i] == needle[0] && n-i >= len(needle) && haystack[i:i+len(needle)] == needle {
            return i
        }
    }
    return -1
}
func main() {
    haystack := "asad"
    needle := "sad"
    fmt.Println(strStr(haystack, needle))
}
