package main

func continuousSubarrays(nums []int) int64 {
	l, r := 0, 0
	cnt := 0
	maxVal := nums[0]
	minVal := nums[0]
	for ; r < len(nums); r++ {
		maxVal = max(maxVal, nums[r])
		minVal = min(minVal, nums[r])
		if maxVal-minVal <= 2 {
			continue
		}
		cnt += countSubArr(r - l)

		maxVal, minVal = nums[r], nums[r]
		for i := r; i >= l; i-- {
			if nums[r]-nums[i] > 2 || nums[r]-nums[i] < -2 {
				l = i + 1
				break
			}
			maxVal = max(maxVal, nums[i])
			minVal = min(minVal, nums[i])
		}

		cnt -= countSubArr(r - l)

	}

	cnt += countSubArr(r - l)
	return int64(cnt)
}

func countSubArr(l int) int {
	return l * (l + 1) / 2
}

func main() {
	println(continuousSubarrays([]int{42, 41, 42, 41, 41, 40, 39, 38}))
}

// 33 34 33 32 31
