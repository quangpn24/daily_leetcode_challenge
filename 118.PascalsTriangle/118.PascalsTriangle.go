package main

import "fmt"

func generate(numRows int) [][]int {
    res := [][]int{[]int{1}}
    for i := 2; i <= numRows; i++ {
        tmp := []int{}
        for j := 0; j < i; j++ {
            if j == 0 || j == i-1 {
                tmp = append(tmp, 1)
            } else {
                tmp = append(tmp, res[i-2][j]+res[i-2][j-1])
            }
        }
        res = append(res, tmp)
    }
    return res
}
func main() {
    fmt.Println(generate(5))
}
