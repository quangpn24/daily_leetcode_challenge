package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, element := range arr1 {
		m[element]++
	}

	for _, e := range arr2 {
		if v, ok := m[e]; ok {
			for i := 0; i < v; i++ {
				res = append(res, e)
			}
			delete(m, e)
		}
	}

	arrTmp := []int{}
	for key, value := range m {
		for i := 0; i < value; i++ {
			arrTmp = append(arrTmp, key)
		}
	}
	sort.Ints(arrTmp)
	return append(res, arrTmp...)
}
func main() {

}
