package main

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	m := make(map[int][]int)

	for i := 0; i < n; i++ {
		m[i] = []int{i + 1}
	}

	res := make([]int, len(queries))
	dist := make([]int, n)
	for i := range dist {
		dist[i] = i // Distance from city 0 to city i
	}

	for i, q := range queries {
		v := m[q[0]]
		v = append(v, q[1])
		m[q[0]] = v
		bfs(m, q[0], n-1, &dist)
		res[i] = dist[n-1]
	}

	return res
}

func bfs(m map[int][]int, start int, end int, dist *[]int) {
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		c := queue[0]
		for _, city := range m[c] {
			if (*dist)[city] > (*dist)[c]+1 {
				(*dist)[city] = (*dist)[c] + 1
				queue = append(queue, city)
			}
			if city == end {
				return
			}
		}
		queue = queue[1:]
	}
}
func main() {
	res := shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}})
	for _, v := range res {
		print(v, " ")
	}
}
