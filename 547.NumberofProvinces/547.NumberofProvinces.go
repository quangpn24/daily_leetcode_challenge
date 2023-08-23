package main

import "fmt"

func DFS(isConnected [][]int, visited []bool, point int) {
	visited[point] = true
	n := len(isConnected)
	for i := 0; i < n; i++ {
		if isConnected[point][i] == 1 && !visited[i] {
			DFS(isConnected, visited[:], i)
		}
	}
}
func findCircleNum(isConnected [][]int) int {
	count := 0
	n := len(isConnected)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			DFS(isConnected, visited[:], i)
			count++
		}
	}
	return count
}
func main() {
	isConnected := [][]int{[]int{1, 0, 0, 1}, []int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 0, 1, 1}}
	fmt.Println(findCircleNum(isConnected))
}
