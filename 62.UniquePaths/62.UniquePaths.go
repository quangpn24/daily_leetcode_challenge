package main

import "fmt"

func uniquePaths(m int, n int) int {
    res := [][]int{}
    for i := 0; i <= m; i++ {
        res = append(res, make([]int, n+1))
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if i == m-1 && j == n-1 {
                res[i][j] = 1
                continue
            }
            res[i][j] = res[i+1][j] + res[i][j+1]
        }
    }
    return res[0][0]
}
func main() {
    fmt.Println(uniquePaths(1, 1))
}
