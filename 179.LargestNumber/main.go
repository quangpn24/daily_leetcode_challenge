package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	arr := make([]string, len(nums))

	for i, n := range nums {
		arr[i] = strconv.Itoa(n)
	}

	sort.Slice(arr, func(i, j int) bool {
		n1, _ := strconv.Atoi(arr[i] + arr[j])
		n2, _ := strconv.Atoi(arr[j] + arr[i])

		if n1 > n2 {
			return true
		}
		return false
	})

	str := strings.Join(arr, "")
	i := 0

	for str[i] == '0' && i < len(str)-1 {
		i++
	}

	return str[i:]
}
func main() {
	println(largestNumber([]int{0, 0}))
}
