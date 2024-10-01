package main

func maxSlidingWindow(nums []int, k int) []int {
	// using monotonic queue
	deque := []int{} // store index, decreasing
	res := make([]int, 0)
	start, end := 0, 0
	for end < len(nums) {
		if len(deque) > 0 && deque[0] < start {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[end] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, end)
		if end-start+1 == k {
			res = append(res, nums[deque[0]])
			start++
		}
		end++
	}

	return res
}

func main() {
	res := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	for _, re := range res {
		println(re, " ")
	}
}
