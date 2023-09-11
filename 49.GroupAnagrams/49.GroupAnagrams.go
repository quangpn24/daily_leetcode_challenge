package main

import (
    "fmt"
    "sort"
)

func parse(str string) string {
    b := []byte(str)
    sort.Slice(b, func(i, j int) bool {
        return b[i] < b[j]
    })
    return string(b)
}
func groupAnagrams(strs []string) [][]string {
    mGroup := make(map[string][]string)
    for _, str := range strs {
        key := parse(str)
        mGroup[key] = append(mGroup[key], str)
    }
    res := [][]string{}
    for _, value := range mGroup {
        res = append(res, value)
    }
    return res
}
func main() {
    strs := []string{"and", "dan"}
    fmt.Println(groupAnagrams(strs))
}
