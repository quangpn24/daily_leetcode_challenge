package main

func shortestSubarray(nums []int, k int) int {
	prefix := make([]int, len(nums)+1)

	prefix[0] = 0
	for i := 1; i <= len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	end := 0
	dequeue := make([]int, 0)

	res := len(nums) + 1
	for end < len(prefix) {
		for len(dequeue) > 0 && prefix[end]-prefix[dequeue[len(dequeue)-1]] >= k {
			res = min(res, end-dequeue[0])
			dequeue = dequeue[1:]
		}

		for len(dequeue) > 0 && prefix[end] <= prefix[dequeue[len(dequeue)-1]] {
			dequeue = dequeue[:len(dequeue)-1]
		}

		dequeue = append(dequeue, end)
		end++
	}

	if res == len(nums)+1 {
		res = -1
	}

	return res
}

func main() {
	println(shortestSubarray([]int{2, -1, 2}, 3))
}
