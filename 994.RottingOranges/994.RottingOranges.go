package main

import "fmt"

func orangesRotting(grid [][]int) int {
	dX := []int{0, -1, 0, 1}
	dY := []int{-1, 0, 1, 0}
	queue := make([][]int, 0)
	row, col := len(grid), len(grid[0])
	count := row * col
	for i, rows := range grid {
		for j, cell := range rows {
			if cell == 0 {
				count--
			} else if cell == 2 {
				count--
				queue = append(queue, []int{i, j})
			}
		}
	}
	if count == 0 {
		return 0
	}
	step := 0
	for len(queue) > 0 {
		for _, rottenOrange := range queue {
			x, y := rottenOrange[0], rottenOrange[1]
			for k := 0; k < 4; k++ {
				cellX := x + dX[k]
				cellY := y + dY[k]
				if cellX >= 0 && cellX <= row-1 && cellY >= 0 && cellY <= col-1 && grid[cellX][cellY] == 1 {
					queue = append(queue, []int{cellX, cellY})
					grid[cellX][cellY] = 2
					count--
				}
			}
			queue = queue[1:]
		}
		if len(queue) > 0 {
			step++
		}
	}
	if count > 0 {
		return -1
	}
	return step
}
func main() {
	grid := [][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}}
	fmt.Println(orangesRotting(grid))
}
