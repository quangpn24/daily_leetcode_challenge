package main

func jump(nums []int) int {
	var l, r, res int

	for r < len(nums)-1 {
		farthest := 0
		for i := l; i < r+1; i++ {
			if i+nums[i] > farthest {
				farthest = i + nums[i]
			}
		}
		l = r + 1
		r = farthest
		res += 1
	}

	return res
}
func main() {
	//println(jump([]int{1, 1, 1, 1}))
	println(jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
	//println(jump([]int{2, 3, 1}))
	//println(jump([]int{1, 2, 3, 4}))
}
