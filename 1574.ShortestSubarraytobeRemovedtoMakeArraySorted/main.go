package main

func findLengthOfShortestSubarray(arr []int) int {
	// prefix arr non-decreasing
	n := len(arr)
	l := 0
	for l < n-1 && arr[l] <= arr[l+1] {
		l++
	}

	// if arr sorted
	if l == n-1 {
		return 0
	}

	// suffix arr non-decreasing
	r := n - 1
	for r > 0 && arr[r] >= arr[r-1] {
		r--
	}

	res := min(n-1-l, r)

	i, j := 0, r
	for i <= l && j < n {
		if arr[i] <= arr[j] {
			res = min(res, j-i-1)
			i++
		} else {
			j++
		}
	}
	return res
}
func main() {
	println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 0, 7, 8, 9}))
}
