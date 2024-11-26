package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxMatrixSum(matrix [][]int) int64 {
	sum := 0
	numOfNegativeNumber := 0
	haveZero := false
	minValue := abs(matrix[0][0])

	for _, row := range matrix {
		for _, cell := range row {
			if cell == 0 {
				haveZero = true
			} else if cell < 0 {
				numOfNegativeNumber++
			}
			sum += abs(cell)
			minValue = min(minValue, abs(cell))
		}
	}

	if haveZero || numOfNegativeNumber%2 == 0 {
		return int64(sum)
	}

	return int64(sum - 2*abs(minValue))
}

func main() {
	println(maxMatrixSum([][]int{{-8, -9, -10, 1}, {-5, -7, -10, -3}, {-8, -4, -2, -6}, {2, -1, 8, -3}}))
}
