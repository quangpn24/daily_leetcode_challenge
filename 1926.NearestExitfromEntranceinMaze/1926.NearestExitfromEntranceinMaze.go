package main

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	near := [][]int{[]int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1}}

	row, col := len(maze), len(maze[0])
	queue := [][]int{entrance}
	count := 0
	length := 1
	i := 1
	maze[entrance[0]][entrance[1]] = '+'
	for ; i <= length; i++ {
		x, y := queue[0][0], queue[0][1]
		if (x != entrance[0] || y != entrance[1]) && (x == 0 || y == 0 || x == row-1 || y == col-1) {
			return count
		}
		for _, p := range near {
			cX := x + p[0]
			cY := y + p[1]
			if cX >= 0 && cX < row && cY >= 0 && cY < col && maze[cX][cY] == '.' {
				maze[cX][cY] = '+'
				queue = append(queue, []int{cX, cY})
			}
		}
		queue = queue[1:]
		if i == length {
			count++
			length += len(queue)
		}
	}
	return -1
}
func main() {
	maze := [][]byte{[]byte{'+', '+', '+'}, []byte{'.', '.', '.'}, []byte{'+', '+', '+'}}
	entrance := []int{1, 0}
	fmt.Println(nearestExit(maze, entrance))
}
