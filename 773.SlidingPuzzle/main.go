package main

type data struct {
	state   string
	step    int
	zeroPos int
}

func slidingPuzzle(board [][]int) int {
	target := "123450"
	queue := make([]data, 0)
	state, zeroPos := boardToString(board)
	queue = append(queue, data{state, 0, zeroPos})

	visited := make(map[string]bool)

	moves := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1, 5},
		3: {0, 4},
		4: {1, 3, 5},
		5: {2, 4},
	}

	for len(queue) > 0 {
		curr := queue[0]

		if curr.state == target {
			return curr.step
		}

		for _, neighbor := range moves[curr.zeroPos] {
			//copy curr state
			newState := []rune(curr.state)
			newState[curr.zeroPos], newState[neighbor] = newState[neighbor], newState[curr.zeroPos]

			if !visited[string(newState)] {
				visited[string(newState)] = true
				queue = append(queue, data{string(newState), curr.step + 1, neighbor})
			}
		}

		queue = queue[1:]
	}
	return -1
}
func boardToString(board [][]int) (string, int) {
	res := ""
	zeroPos := -1
	for _, row := range board {
		for _, v := range row {
			if v == 0 {
				zeroPos = len(res)
			}
			res += string(v + '0')
		}
	}
	return res, zeroPos
}
func main() {
	println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
}
