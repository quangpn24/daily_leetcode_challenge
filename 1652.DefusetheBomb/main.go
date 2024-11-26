package main

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	n := len(code)
	if k == 0 {
		return res
	}

	if k > 0 {
		l, r := 0, 1
		sum := 0
		for l < n {
			sum += code[r%n]
			if r-l == k {
				res[l] = sum
				l++
				sum -= code[l%n]
			}
			r++
		}
	} else {
		sum := 0
		for i := n - 1; i >= n+k; i-- {
			sum += code[i]
		}
		res[0] = sum
		l, r := 0, n+k
		for l < n-1 {
			sum = sum + code[l] - code[r%n]
			l++
			r++
			res[l] = sum
		}
	}
	return res
}
func main() {
	res := decrypt([]int{5, 7, 1, 4}, 3)
	for _, v := range res {
		print(v, " ")
	}
}
