package main

func jump(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				if arr[i] == 0 {
					arr[i] = arr[j] + 1
				} else {
					arr[i] = min(arr[j]+1, arr[i])
				}
			}
		}
	}
	return arr[len(arr)-1]
}
func main() {
	//println(jump([]int{1, 1, 1, 1}))
	println(jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
	//println(jump([]int{2, 3, 1}))
	//println(jump([]int{1, 2, 3, 4}))
}
