package main

import "fmt"

func search(nums []int, target int) int {
	if nums[0] < nums[len(nums)-1] {
		return binarySearch(nums, target)
	}

	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if nums[0] == target {
		return 0
	} else if nums[0] < target {
		return binarySearch(nums[:r], target)
	}

	res := binarySearch(nums[r:], target)
	if res != -1 {
		return res + r
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func main() {
	fmt.Println(search([]int{1, 3}, 3))
}
