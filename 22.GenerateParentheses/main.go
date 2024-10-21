package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(res *[]string, str string, open, close, n int) {
	if open < close {
		return
	}
	if open == n && close == n {
		*res = append(*res, str)
		return
	}

	if open < n {
		str += "("
		open++
		backtrack(res, str, open, close, n)
		open--
		str = str[:len(str)-1]
	}

	if close < n {
		str += ")"
		close++
		backtrack(res, str, open, close, n)
		close--
		str = str[:len(str)-1]
	}
}
func main() {
	res := generateParenthesis(1)
	for _, re := range res {
		println(re)
	}
}
