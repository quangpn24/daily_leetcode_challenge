package main

func canSortArray(nums []int) bool {
	start := 0
	end := len(nums) - 1
	for start < end {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] <= nums[i+1] {
				start++
				continue
			}

			if countSetBits(nums[i]) == countSetBits(nums[i+1]) {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				start = 0
			} else {
				return false
			}
		}
	}
	return true
}
func countSetBits(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

func main() {
	println(canSortArray([]int{16, 8, 4, 2}))
}
