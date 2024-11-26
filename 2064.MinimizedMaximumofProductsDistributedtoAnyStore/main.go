package main

func minimizedMaximum(n int, quantities []int) int {
	sum := 0
	for _, quantity := range quantities {
		sum += quantity
	}

	l, r := 0, sum
	for l < r {
		mid := (r + l) / 2
		a := numOfStoreDistributedProducts(mid, quantities)
		if a > n {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if numOfStoreDistributedProducts(l, quantities) <= n {
		return l
	}
	return l + 1
}

func numOfStoreDistributedProducts(maxProduct int, quantities []int) int {
	cnt := 0
	for _, quantity := range quantities {
		cnt += quantity / maxProduct
		if quantity%maxProduct != 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	println(minimizedMaximum(3, []int{2, 10, 6}))
	//println(numOfStoreDistributedProducts(8, ))
}
