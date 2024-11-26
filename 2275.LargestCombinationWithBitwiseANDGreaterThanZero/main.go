package main

func largestCombination(candidates []int) int {
	cnt := 0
	for _, candidate := range candidates {
		if (candidate & 1) == 1 {
			cnt++
			println(candidate)
		}
	}
	return cnt
}

func main() {
	//println(largestCombination([]int{16, 17, 71, 62, 12, 24, 14}))
	println(4 >> 2)
}
