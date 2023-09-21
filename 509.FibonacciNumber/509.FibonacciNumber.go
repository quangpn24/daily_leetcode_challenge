package main

import "fmt"

func fib(n int) int {
    if n == 0 {
        return 0
    }
    ans := make([]int, n+1)
    ans[0] = 0
    ans[1] = 1
    for i := 2; i <= n; i++ {
        ans[i] = ans[i-1] + ans[i-2]
    }
    return ans[n]
}
func main() {
    fmt.Println(fib(3))
}
