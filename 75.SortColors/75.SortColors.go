package main

import "fmt"

func sortColors(nums []int) {
	pos0, pos2 := 0, len(nums)-1
	for i := 0; i <= pos2 && pos0 < pos2; {
		if nums[i] == 0 {
			nums[i], nums[pos0] = nums[pos0], nums[i]
			pos0++
			i++
		} else if nums[i] == 2 {
			for nums[pos2] == 2 && i < pos2 {
				pos2--
			}
			nums[i], nums[pos2] = nums[pos2], nums[i]
			pos2--
		} else {
			i++
		}
	}
}
func main() {
	nums := []int{2, 1}
	sortColors(nums)
	fmt.Println(nums)
}
