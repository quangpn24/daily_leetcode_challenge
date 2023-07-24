package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		x = -x
		isNegative = true
	}
	str := strconv.Itoa(x)
	arr := []int{0}
	for _, item := range str {
		num, _ := strconv.Atoi(string(item))
		arr = append(arr, num)
	}

	result := 0
	for i, num := range arr {
		result += num * int(math.Pow10(i-1))
	}
	if isNegative {
		result = -result
	}
	if float64(result) < -math.Pow(2, 31) || float64(result) >= math.Pow(2, 31) {
		return 0
	}
	return result
}
func main() {
	x := 1534236469
	fmt.Println(reverse(x))
}
