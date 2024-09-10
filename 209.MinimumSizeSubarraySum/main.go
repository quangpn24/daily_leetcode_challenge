package main

func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	sum := 0
	l, r := 0, 0
	for sum >= target || r < len(nums) {
		if sum < target {
			sum += nums[r]
			r++
		} else {
			minLen = min(minLen, r-l)
			sum -= nums[l]
			l++
		}
	}

	if minLen < len(nums)+1 {
		return minLen
	}
	return 0
}

func main() {
	println(minSubArrayLen(4, []int{1, 4, 4}))
}
