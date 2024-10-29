package main

//func findTargetSumWays(nums []int, target int) int {
//	cnt := 0
//	backtrack(&cnt, nums, target, 0, 0)
//	return cnt
//}
//
//func backtrack(cnt *int, nums []int, target int, step int, sum int) {
//	if step == len(nums) {
//		if sum == target {
//			*cnt += 1
//		}
//		return
//	}
//	backtrack(cnt, nums, target, step+1, sum+nums[step])
//	backtrack(cnt, nums, target, step+1, sum-nums[step])
//}
func findTargetSumWays(nums []int, target int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}

	if target < 0 {
		target = -target
	}
	
	if sum < target || (sum+target)%2 != 0 {
		return 0
	}

	n := (sum + target) / 2

	dp := make([]int, n+1)
	dp[0] = 1

	for _, num := range nums {
		for i := n; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[n]
}

func main() {
	println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
