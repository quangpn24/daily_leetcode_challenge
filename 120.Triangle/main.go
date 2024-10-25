package main

func minimumTotal(triangle [][]int) int {
	arr := make([][]int, len(triangle))
	arr[0] = []int{triangle[0][0]}
	for i := 1; i < len(triangle); i++ {
		tmp := make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			pre, curr := 10000, 10000
			if j-1 >= 0 {
				pre = arr[i-1][j-1]
			}

			if j < len(arr[i-1]) {
				curr = arr[i-1][j]
			}

			tmp[j] = triangle[i][j] + min(pre, curr)
		}
		arr[i] = tmp
	}

	minSumPath := 10000
	for _, s := range arr[len(arr)-1] {
		minSumPath = min(minSumPath, s)
	}
	return minSumPath
}

func main() {
	println(minimumTotal([][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}))
}
