package main

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	res := make([]int, 3)
	n := len(s)
	for i := 0; i < n; i++ {
		res[s[i]-'a']++
	}

	if res[0] < k || res[1] < k || res[2] < k {
		return -1
	}

	l, r := -1, n-1

	minMinute := n + 1
	res = []int{0, 0, 0}
	for l < n {
		if l > -1 {
			res[s[l]-'a']++
		}
		if isOk(res, k) {
			for isOk(res, k) && r < n {
				res[s[r]-'a']--
				r++
			}
			if !isOk(res, k) {
				r--
				res[s[r]-'a']++
			}
		} else {
			for !isOk(res, k) && l < r {
				res[s[r]-'a']++
				r--
			}
			r++
		}
		minMinute = min(minMinute, res[0]+res[1]+res[2])
		l++
	}
	return minMinute
}

func isOk(res []int, k int) bool {
	return res[0] >= k && res[1] >= k && res[2] >= k
}

func main() {
	println(takeCharacters("cbbac", 1))
}
