package main

import "fmt"

func titleToNumber(columnTitle string) int {
    ans := 0
    for _, c := range columnTitle {
        ans *= 26
        ans += int(c) - 64
    }
    return ans
}
func main() {
    fmt.Println(titleToNumber("AA"))
}
