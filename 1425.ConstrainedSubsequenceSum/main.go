package main

func constrainedSubsetSum(nums []int, k int) int {
	end := 0
	maxSum := 0
	dequeue := make([]int, 0)
	for end < len(nums) {
		
		dequeue = append(dequeue, end)
		maxSum += nums[end]
	}
	return maxSum
}
func main() {

}
