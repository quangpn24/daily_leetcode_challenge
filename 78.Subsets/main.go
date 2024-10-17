package main

func subsets(nums []int) [][]int {
	var res = make([][]int, 0)
	var arr = make([]int, 0)
	res = append(res, []int{})
	try(0, nums, &res, arr)
	return res
}

func try(i int, nums []int, res *[][]int, arr []int) {
	for j := i; j < len(nums); j++ {
		arr = append(arr, nums[j])
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
		try(j+1, nums, res, arr)
		arr = arr[:len(arr)-1]
	}
}

func main() {
	result := subsets([]int{1, 2, 3})
	for _, ints := range result {
		for _, i := range ints {
			print(i, " ")
		}
		println("")
	}
}
