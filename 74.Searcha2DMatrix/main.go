package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	colLo, colHi := 0, m-1
	for colLo <= colHi {
		mid := colLo + (colHi-colLo)/2
		if matrix[mid][n-1] < target {
			colLo = mid + 1
		} else if matrix[mid][n-1] > target {
			if matrix[mid][0] <= target {
				return binarySearch(matrix[mid], target)
			} else {
				colHi = mid - 1
			}
		} else {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
func main() {
	matrix := [][]int{
		{1},
	}
	println(searchMatrix(matrix, 1))
}
