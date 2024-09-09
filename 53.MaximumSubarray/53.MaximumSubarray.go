package main

func maxSubArray(nums []int) int {
	res := nums[0]
	arr := make([]int, len(nums)+1)
	for i, num := range nums {
		arr[i+1] = max(arr[i]+num, num)

		res = max(res, arr[i+1])
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	println(maxSubArray(nums))
}
