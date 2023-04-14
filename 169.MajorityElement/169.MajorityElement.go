package main

func majorityElement(nums []int) int {
	myMap := map[int]int{}
	for _, num := range nums {
		if v, ok := myMap[num]; ok {
			myMap[num] = v + 1
		} else {
			myMap[num] = 1
		}
	}
	max := 0
	res := nums[0]
	for key, val := range myMap {
		if max < val {
			max = val
			res = key
		}
	}
	return res
}
