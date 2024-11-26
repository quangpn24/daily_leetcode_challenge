package main

func maxMoves(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	matrix := make([][]int, len(grid))
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
	}

	canMove := false

	for i := 0; i < row; i++ {
		if grid[i][0] < grid[i][1] {
			canMove = true
			matrix[i][1] = 1
		}

		if i > 0 && grid[i][0] < grid[i-1][1] {
			canMove = true
			matrix[i-1][1] = 1
		}

		if i < row-1 && grid[i][0] < grid[i+1][1] {
			canMove = true
			matrix[i+1][1] = 1
		}
	}

	if !canMove {
		return 0
	}

	res := 1

	for i := 1; i < col-1; i++ {
		for j := 0; j < row; j++ {
			if matrix[j][i] == 0 {
				continue
			}

			if grid[j][i] < grid[j][i+1] {
				canMove = true
				matrix[j][i+1] = max(matrix[j][i]+1, matrix[j][i+1])
				res = max(res, matrix[j][i+1])
			}

			if j > 0 && grid[j][i] < grid[j-1][i+1] {
				canMove = true
				matrix[j-1][i+1] = max(matrix[j][i]+1, matrix[j-1][i+1])
				res = max(res, matrix[j-1][i+1])
			}

			if j < row-1 && grid[j][i] < grid[j+1][i+1] {
				canMove = true
				matrix[j+1][i+1] = max(matrix[j][i]+1, matrix[j+1][i+1])
				res = max(res, matrix[j+1][i+1])
			}
		}
	}

	return res
}
func main() {
	println(maxMoves([][]int{
		[]int{187, 167, 209, 251, 152, 236, 263, 128, 135},
		[]int{267, 249, 251, 285, 73, 204, 70, 207, 74},
		[]int{189, 159, 235, 66, 84, 89, 153, 111, 189},
		[]int{120, 81, 210, 7, 2, 231, 92, 128, 218},
		[]int{193, 131, 244, 293, 284, 175, 226, 205, 245}}))
}
