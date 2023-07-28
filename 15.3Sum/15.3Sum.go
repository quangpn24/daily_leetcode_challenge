package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func convertIntsToString(nums []int) string {
	sort.Ints(nums)
	res := ""
	for i := 0; i < len(nums)-1; i++ {
		res += strconv.Itoa(nums[i]) + "|"
	}
	res += strconv.Itoa(nums[len(nums)-1])
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	myMap := map[int][][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			value := nums[i] + nums[j]
			if value+min <= 0 && value+max >= 0 {
				myMap[value] = append(myMap[value], []int{i, j})
			}
		}
	}

	resMap := map[string]bool{}
	for k, num := range nums {
		if values, ok := myMap[-num]; ok {
			for _, arr := range values {
				if arr[0] != k && arr[1] != k {
					str := convertIntsToString([]int{nums[arr[0]], nums[arr[1]], num})
					if _, ok := resMap[str]; !ok {
						resMap[str] = true
					}
				}
			}
		}
	}

	res := [][]int{}
	for key := range resMap {
		tmps := strings.Split(key, "|")
		numi, _ := strconv.Atoi(tmps[0])
		numj, _ := strconv.Atoi(tmps[1])
		numk, _ := strconv.Atoi(tmps[2])
		res = append(res, []int{numi, numj, numk})
	}
	return res
}
func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}
