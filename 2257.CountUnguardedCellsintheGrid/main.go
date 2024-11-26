package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = 2
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 2
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	for i := 0; i < len(guards); i++ {
		for j := 0; j < 4; j++ {
			x := guards[i][0] + dx[j]
			y := guards[i][1] + dy[j]
			for x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != 2 {
				grid[x][y] = 1
				x += dx[j]
				y += dy[j]
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
			}
		}
	}

	return res
}
func main() {
	println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}
