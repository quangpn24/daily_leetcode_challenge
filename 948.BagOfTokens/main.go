package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	l, r := 0, len(tokens)-1
	score := 0
	maxScore := 0
	for l <= r {
		if power >= tokens[l] {
			score++
			maxScore = max(maxScore, score)
			power -= tokens[l]
			l++
		} else if score > 0 {
			score--
			power += tokens[r]
			r--
		} else {
			break
		}
	}

	return maxScore
}
func main() {
	println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
