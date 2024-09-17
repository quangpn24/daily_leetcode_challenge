package main

func minDays(bloomDay []int, m int, k int) int {
	maxD := 0
	for _, d := range bloomDay {
		maxD = max(maxD, d)
	}
	l, r := 1, maxD

	for l < r {
		mid := l + (r-l)/2
		n := calcBouquets(bloomDay, mid, k)
		if n < m {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if calcBouquets(bloomDay, l, k) < m {
		return -1
	}

	return l
}

func calcBouquets(bloomDay []int, day int, k int) int {
	cnt := 0
	nDay := 0
	for _, d := range bloomDay {
		if d <= day {
			nDay++
		} else {
			cnt += nDay / k
			nDay = 0
		}
	}
	return cnt + nDay/k
}

func main() {
	bloomDay := []int{1000000000, 1000000000}
	println(minDays(bloomDay, 1, 1))
}
