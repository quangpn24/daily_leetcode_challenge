package main

func minimumCardPickup(cards []int) int {
	res := 1000001
	start, end := 0, 0
	m := make(map[int]int)

	for end < len(cards) {
		m[cards[end]]++

		for end-start == len(m) {
			res = min(res, end-start+1)
			m[cards[start]]--
			if m[cards[start]] == 0 {
				delete(m, cards[start])
			}
			start++
		}
		end++
	}

	if res == 1000001 {
		return -1
	}

	return res
}
func main() {
	println(minimumCardPickup([]int{95, 11, 8, 65, 5, 86, 30, 27, 30, 73, 15, 91, 30, 7, 37, 26, 55, 76, 60, 43, 36, 85, 47, 96, 6}))
}
