package main

import "fmt"

func climbStairs(n int) int {
	var arrRes = []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		arrRes = append(arrRes, arrRes[i-2]+arrRes[i-1])
	}
	return arrRes[n]
}

func main() {
	fmt.Println(climbStairs(5))
}
