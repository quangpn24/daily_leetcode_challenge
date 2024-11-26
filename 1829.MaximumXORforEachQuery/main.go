package main

func getMaximumXor(nums []int, maximumBit int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, len(nums))
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
		res[len(nums)-i-1] = xor ^ (1<<maximumBit - 1)
	}

	return res
}

func main() {
	res := getMaximumXor([]int{2, 3, 4, 7}, 3)
	for _, re := range res {
		print(re, " ")
	}
}
