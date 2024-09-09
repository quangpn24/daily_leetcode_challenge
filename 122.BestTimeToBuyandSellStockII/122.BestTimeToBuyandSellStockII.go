package main

// if prices[i] > prices[i-1], then profit += prices[i] - prices[i-1]
// It's mean that we can buy and sell at the same day.
func maxProfit(prices []int) int {
	profit := 0
	buyPrice := prices[0]
	for _, p := range prices {
		if p > buyPrice {
			profit += p - buyPrice
		}
		buyPrice = p
	}
	return profit
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	println(ans)
}
