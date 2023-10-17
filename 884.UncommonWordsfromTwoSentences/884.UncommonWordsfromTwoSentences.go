package main

import (
    "fmt"
    "strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
    res := []string{}
    arr1 := strings.Split(s1, " ")
    arr2 := strings.Split(s2, " ")
    m := make(map[string]int)
    for _, s := range arr1 {
        m[s]++
    }
    for _, s := range arr2 {
        m[s]++
    }
    for k, v := range m {
        if v == 1 {
            res = append(res, k)
        }
    }
    return res
}
func main() {
    fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}
