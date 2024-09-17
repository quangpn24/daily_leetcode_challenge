package main

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)/2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func searchRange(nums []int, target int) []int {
	return []int{lowerBound(nums, target), upperBound(nums, target)}
}
func main() {
	//println(lowerBound([]int{5, 7, 7, 8, 8, 10}, 8))
	println(upperBound([]int{5, 7, 7, 8, 8, 10}, 8))
}
