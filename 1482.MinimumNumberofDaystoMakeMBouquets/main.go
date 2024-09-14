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
	i := 0
	for i <= len(bloomDay)-k {
		if bloomDay[i] <= day && isMakeBouquets(bloomDay, k, i, day) {
			cnt++
			i += k - 1
		}
		i++
	}
	return cnt
}

func isMakeBouquets(bloomDay []int, k int, start int, day int) bool {
	for i := start; i < start+k; i++ {
		if bloomDay[i] > day {
			return false
		}
	}
	return true
}

func main() {
	bloomDay := []int{1000000000, 1000000000}
	println(minDays(bloomDay, 1, 1))
}
