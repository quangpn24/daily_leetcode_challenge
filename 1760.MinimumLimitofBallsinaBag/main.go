package main

func countOperation(nums []int, maxBall int) int {
	cnt := 0
	for _, num := range nums {
		cnt += (num - 1) / maxBall
	}
	return cnt
}

func minimumSize(nums []int, maxOperations int) int {
	maxBall := 0
	for _, num := range nums {
		maxBall = max(maxBall, num)
	}

	l, r := 0, maxBall

	for l < r {
		mid := l + (r-l)/2
		a := countOperation(nums, mid)
		if a > maxOperations {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if countOperation(nums, l) == maxOperations {
		return l
	}

	return l + 1
}

func main() {
	//println(minimumSize([]int{2, 4, 8, 2}, 4))
	println(minimumSize([]int{9}, 2))
	//println(countOperation([]int{9}, 2))
}
