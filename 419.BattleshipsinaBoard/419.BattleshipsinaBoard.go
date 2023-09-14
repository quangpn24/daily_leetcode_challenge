package main

import "fmt"

var dX = []int{-1, 0, 1, 0}
var dY = []int{0, -1, 0, 1}

func countBattleships(board [][]byte) int {
    m := len(board)
    n := len(board[0])
    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'X' {
                DFS(board, []int{i, j}, m, n)
                cnt++
            }
        }
    }
    return cnt
}
func DFS(board [][]byte, pos []int, m int, n int) {
    for k := 0; k < 4; k++ {
        nextX := pos[0] + dX[k]
        nextY := pos[1] + dY[k]
        if nextX >= 0 && nextX < m && nextY >= 0 && nextY < n && board[nextX][nextY] == 'X' {
            board[nextX][nextY] = '.'
            DFS(board, []int{nextX, nextY}, m, n)
        }
    }
}

func main() {
    board := [][]byte{[]byte{'X', '.', '.', 'X'}, []byte{'.', '.', '.', 'X'}, []byte{'.', '.', '.', 'X'}}
    fmt.Println(countBattleships(board))
}
