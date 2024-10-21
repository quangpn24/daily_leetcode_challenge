package main

func lengthOfLIS(nums []int) int {
	counts := make(map[int]int)
	counts[1] = nums[0]

	res := 0

	for _, num := range nums {
		m := 0
		for k, v := range counts {
			if v < num {
				m = max(m, k)
			}
		}
		res = max(res, m+1)
		if v, ok := counts[m+1]; !ok {
			counts[m+1] = num
		} else {
			counts[m+1] = min(v, num)
		}
	}
	return res
}
func main() {

}
