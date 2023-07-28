package main

import (
	"bytes"
	"fmt"
	"math"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	step := numRows*2 - 2

	//create matrix
	matrix := make([][]uint8, numRows)
	for i := 0; i < len(matrix); i++ {
		length := int(math.Ceil(float64(len(s))/float64(step))) * (numRows - 1)
		matrix[i] = make([]uint8, length)
	}

	index := 0
	row := 0
	col := 0
	for {
		if index < len(s) {
			matrix[row][col] = s[index]
			index++
		} else {
			break
		}

		if index%step < numRows && index%step > 0 {
			row++
		} else {
			col++
			row--
		}
	}

	//result
	res := make([]byte, 0)
	for i := 0; i < len(matrix); i++ {
		res = append(res, matrix[i]...)
	}
	res = bytes.ReplaceAll(res, []byte{0}, []byte{})
	return string(res)
}

func main() {
	str := "PAHNAPLSIIGYIR"
	abc := convert(str, 3)
	fmt.Println(abc)
}
