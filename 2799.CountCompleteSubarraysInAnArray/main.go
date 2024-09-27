package main

func countCompleteSubarrays(nums []int) int {
	mTmp := make(map[int]int)
	for _, num := range nums {
		mTmp[num]++
	}
	numDistinct := len(mTmp)

	cnt := 0
	m := make(map[int]int)
	start, end := 0, 0
	for end < len(nums) {
		m[nums[end]]++
		for len(m) == numDistinct {
			m[nums[start]]--
			if m[nums[start]] == 0 {
				delete(m, nums[start])
			}
			cnt += len(nums) - end
			start++
		}
		end++
	}

	return cnt
}
func main() {
	println(countCompleteSubarrays([]int{1, 3, 1, 2, 2}))
}
