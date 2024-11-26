package main

func minimumSubarrayLength(nums []int, k int) int {
	r := 0
	resOR := 0
	cnt := len(nums) + 1
	for r < len(nums) {
		if nums[r] >= k {
			return 1
		}
		resOR = resOR | nums[r]
		if resOR >= k {
			tmp := nums[r]
			l := r
			for tmp < k {
				l--
				if tmp|nums[l] < k {
					tmp = tmp | nums[l]
				} else {
					break
				}
			}
			cnt = min(cnt, r-l+1)
			resOR = tmp
		}
		r++
	}

	if cnt == len(nums)+1 {
		return -1
	}
	return cnt
}
func main() {
	println(minimumSubarrayLength([]int{1, 2, 32, 21}, 55))
}
