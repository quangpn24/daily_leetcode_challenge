package main

import "fmt"

var dX = []int{0, -1, 0, 1}
var dY = []int{-1, 0, 1, 0}

func BFS(grid [][]byte, pos []int) {
    row := len(grid)
    col := len(grid[0])
    for k := 0; k < 4; k++ {
        cellX := pos[0] + dX[k]
        cellY := pos[1] + dY[k]
        if cellX >= 0 && cellX <= row-1 && cellY >= 0 && cellY <= col-1 && grid[cellX][cellY] == '1' {
            grid[cellX][cellY] = '0'
            BFS(grid, []int{cellX, cellY})
        }
    }
}
func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                BFS(grid, []int{i, j})
                count++
            }
        }
    }
    return count
}
func main() {
    grid := [][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}}
    fmt.Println(numIslands(grid))
}
