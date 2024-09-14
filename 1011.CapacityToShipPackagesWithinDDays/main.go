package main

func shipWithinDays(weights []int, days int) int {
	sum := 0
	maxW := 0
	for _, w := range weights {
		sum += w
		maxW = max(maxW, w)
	}
	l, r := maxW, sum
	for l < r {
		mid := l + (r-l)/2
		d := calcCapacity(mid, weights)
		if d >= days {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func calcCapacity(c int, weights []int) int {
	cnt := 0
	sum := 0
	for _, w := range weights {
		if sum+w > c {
			cnt++
			sum = 0
		}
		sum += w
	}
	return cnt
}
func main() {
	weights := []int{1, 2, 3, 1, 1}
	//println(calcCapacity(2, weights))
	println(shipWithinDays(weights, 4))
}
