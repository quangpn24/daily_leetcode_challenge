package main

import (
    "fmt"
    "strings"
)

func wordPattern(pattern string, s string) bool {
    m := make(map[byte]string)
    m1 := make(map[string]byte)
    arr := strings.Split(s, " ")
    if len(pattern) != len(arr) {
        return false
    }
    for i, c := range pattern {
        if v, ok := m[byte(c)]; ok {
            if v != arr[i] {
                return false
            }
        } else {
            m[byte(c)] = arr[i]
            if v1, ok1 := m1[arr[i]]; ok1 {
                if v1 != byte(c) {
                    return false
                }
            } else {
                m1[arr[i]] = byte(c)
            }
        }
    }
    return true
}

func main() {
    pt := "abba"
    s := "dog cat cat dog"
    fmt.Println(wordPattern(pt, s))
}
