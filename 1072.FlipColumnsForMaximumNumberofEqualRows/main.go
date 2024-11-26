package main

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := make(map[string]int)

	for _, row := range matrix {
		str := ""
		str2 := ""
		for _, v := range row {
			str += string(v + '0')
			str2 += string(1 - v + '0')
		}

		m[min(str, str2)]++
	}

	maxCount := 0
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}
	return maxCount
}
func main() {
	println(maxEqualRowsAfterFlips([][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}))
}
