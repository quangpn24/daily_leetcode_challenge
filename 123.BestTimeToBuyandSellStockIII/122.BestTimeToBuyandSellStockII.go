package main

func maxInt(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	minLeft := prices[0]
	maxProfitLefts := make([]int, len(prices))
	maxRight := prices[len(prices)-1]
	maxProfitRights := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		pl := prices[i]
		maxProfitLefts[i] = maxProfitLefts[i-1]
		if pl > minLeft {
			maxProfitLefts[i] = maxInt(maxProfitLefts[i-1], pl-minLeft)
		} else {
			minLeft = pl
		}

		pr := prices[len(prices)-1-i]
		maxProfitRights[len(prices)-1-i] = maxProfitRights[len(prices)-i]
		if pr < maxRight {
			maxProfitRights[len(prices)-1-i] = maxInt(maxProfitRights[len(prices)-i], maxRight-pr)
		} else {
			maxRight = pr
		}
	}

	res := 0
	for i := 0; i < len(prices); i++ {
		res = maxInt(res, maxProfitLefts[i]+maxProfitRights[i])
	}
	return res
}
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	ans := maxProfit(prices)
	println(ans)
}
