package main

func maximumUniqueSubarray(nums []int) int {
	start, end := 0, 0
	maxScore := 0
	windowScore := 0

	m := make(map[int]int)
	for end < len(nums) {
		windowScore += nums[end]
		m[nums[end]]++
		if end-start+1 > len(m) {
			windowScore -= nums[start]
			m[nums[start]]--
			if m[nums[start]] == 0 {
				delete(m, nums[start])
			}
			start++
			end++
		} else {
			maxScore = max(maxScore, windowScore)
			end++
		}
	}

	for start < end {
		if end-start == len(m) {
			maxScore = max(maxScore, windowScore)
			break
		} else {
			windowScore -= nums[start]
			m[nums[start]]--
			if m[nums[start]] == 0 {
				delete(m, nums[start])
			}
			start++
		}
	}

	return maxScore
}
func main() {
	println(maximumUniqueSubarray([]int{1, 2, 3, 4, 2, 5, 6, 7, 4}))
	println(maximumUniqueSubarray([]int{1, 2, 3, 4, 2, 5, 6, 7, 4, 9}))
	println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
}
