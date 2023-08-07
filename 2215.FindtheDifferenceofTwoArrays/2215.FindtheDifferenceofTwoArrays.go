package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 0)
	map1 := map[int]bool{}
	for _, n := range nums1 {
		if _, ok := map1[n]; !ok {
			map1[n] = true
		}
	}
	ans2 := make([]int, 0)
	for _, n2 := range nums2 {
		if _, ok := map1[n2]; ok {
			map1[n2] = false
		} else {
			ans2 = append(ans2, n2)
			map1[n2] = false
		}
	}
	ans1 := make([]int, 0)
	for key, value := range map1 {
		if value {
			ans1 = append(ans1, key)
		}
	}
	return append(res, ans1, ans2)
}
func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{2, 4, 6, 6}
	fmt.Println(findDifference(nums1, nums2))

}
