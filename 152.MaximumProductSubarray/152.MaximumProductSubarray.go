package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum := nums[0]
	minNum := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxNum, minNum = minNum, maxNum
		}
		maxNum = max(maxNum*nums[i], nums[i])
		minNum = min(minNum*nums[i], nums[i])
		println("maxNum: ", maxNum, "minNum: ", minNum)
		result = max(result, maxNum)
	}
	return result
}

func main() {
	//nums := []int{2, 0, -2, 4}
	//println(maxProduct(nums))
	//
	//

	fmt.Println(ab(" "))
}

func ab(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 0

	m := make(map[rune]bool)
	l, r := 0, 0
	for r < len(s) {
		if _, ok := m[rune(s[r])]; !ok {
			m[rune(s[r])] = true
			r++
		} else {
			delete(m, rune(s[l]))
			maxLen = max(maxLen, r-l)
			l++
		}
	}

	maxLen = max(maxLen, r-l)

	return maxLen
}
