package main

import "fmt"

func DFS(rooms [][]int, room []int, visited []bool, count *int) {
	for _, key := range room {
		if !visited[key] {
			visited[key] = true
			*count = *count - 1
			DFS(rooms, rooms[key], visited, count)
		}
	}
}
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	count := len(rooms) - 1
	DFS(rooms, rooms[0], visited, &count)
	return count == 0
}
func main() {
	rooms := [][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{}}
	fmt.Println(canVisitAllRooms(rooms))
}
