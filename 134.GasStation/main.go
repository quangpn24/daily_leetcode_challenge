package main

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	idx := 0
	total := 0
	for i, g := range gas {
		sum += g - cost[i]
		total += g - cost[i]
		if sum < 0 {
			idx = i + 1
			sum = 0
		}
	}
	if total < 0 {
		return -1
	}
	return idx
}
func main() {
	println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
