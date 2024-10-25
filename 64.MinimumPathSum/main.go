package main

import "math"

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		arr := make([]int, col)
		matrix[i] = arr
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				arr[j] = grid[i][j]
				continue
			}

			left, above := math.MaxInt, math.MaxInt
			if i-1 >= 0 {
				above = grid[i-1][j]
			}

			if j-1 >= 0 {
				left = grid[i][j-1]
			}

			arr[j] = min(left, above) + grid[i][j]
		}
	}

	return matrix[row-1][col-1]
}

func main() {
	println(minPathSum([][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}))
}
