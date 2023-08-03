package main

import "fmt"

func longestSubarray(nums []int) int {
	maxLength := -1
	zeroIndex := -1
	flag := false
	l := 0
	r := 0
	lenNums := len(nums)
	for r < lenNums {
		if nums[r] == 1 {
			r++
		} else if !flag {
			zeroIndex = r
			r++
			flag = true
		} else {
			if maxLength < r-l {
				maxLength = r - l
			}
			if nums[l] == 0 {
				l++
				zeroIndex = r
			} else {
				l = zeroIndex + 1
				zeroIndex = r
			}
			r++
		}
	}
	if r-l > maxLength {
		maxLength = r - l
	}
	return maxLength - 1
}
func longestSubarray2(nums []int) int {
	maxLength := -1
	count := 0
	res := []int{}
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			res = append(res, count)
			count = 0
		}
	}
	res = append(res, count)
	for i := 0; i < len(res)-1; i++ {
		if res[i]+res[i+1] > maxLength {
			maxLength = res[i] + res[i+1]
		}
	}
	if len(res) == 1 {
		maxLength = res[0] - 1
	}
	return maxLength
}
func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray2(nums))
}
