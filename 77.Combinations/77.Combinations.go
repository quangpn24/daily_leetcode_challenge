package main

import "fmt"

var result = [][]int{}
var res = []int{}

func combine(n int, k int) [][]int {
    result = [][]int{}
    res = []int{}
    choose(1, n, k)
    return result
}

func choose(i int, n int, k int) {
    for j := i; j <= n; j++ {
        res = append(res, j)
        if len(res) < k {
            choose(j+1, n, k)
        } else {
            dst := make([]int, k)
            copy(dst, res)
            result = append(result, dst)
            res = res[:k-1]
        }
    }
    if len(res) > 0 {
        res = res[:len(res)-1]
    }
}
func main() {
    fmt.Println(combine(1, 1))
}
