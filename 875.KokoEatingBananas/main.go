package main

func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		maxPile = max(maxPile, p)
	}
	l, r := 1, maxPile
	for l < r {
		mid := l + (r-l)/2
		t := calcHours(piles, mid)
		if t > h {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func calcHours(piles []int, k int) int {
	index := 0
	h := 0
	for index < len(piles) {
		if piles[index]%k == 0 {
			h += piles[index] / k
		} else {
			h += piles[index]/k + 1
		}
		index++
	}
	return h
}

func main() {
	println(minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316,
		877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 823855818))
}
