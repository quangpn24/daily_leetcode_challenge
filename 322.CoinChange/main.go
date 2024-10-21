package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	mem := make(map[int]int)
	for _, c := range coins {
		mem[c] = 1
	}
	x := dp(coins, amount, &mem)
	return x
}

func dp(coins []int, n int, mem *map[int]int) int {
	if (*mem)[n] != 0 {
		return (*mem)[n]
	}

	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] > n {
			continue
		}
		min1 := dp(coins, n-coins[i], mem)

		if min1 != -1 {
			if (*mem)[n] == 0 {
				(*mem)[n] = min1 + 1
			} else {
				(*mem)[n] = min((*mem)[n], min1+1)
			}
		}
	}
	if (*mem)[n] == 0 {
		(*mem)[n] = -1
	}
	return (*mem)[n]
}

func main() {
	println(coinChange([]int{1, 2, 5}, 11))
}
