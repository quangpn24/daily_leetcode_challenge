package main

func maxDistance(arrays [][]int) int {
	largestNumber := -100000
	largestIndex := -1
	secondLargestNumber := -100000

	smallestNumber := 100000
	smallestIndex := -1
	secondSmallestNumber := 100000

	for i, array := range arrays {
		n := array[len(array)-1]
		n0 := array[0]
		if largestNumber <= n {
			secondLargestNumber = largestNumber
			largestNumber = n
			largestIndex = i
		} else if secondLargestNumber < n {
			secondLargestNumber = n
		}

		if smallestNumber >= n0 {
			secondSmallestNumber = smallestNumber
			smallestIndex = i
			smallestNumber = n0
		} else if secondSmallestNumber > n0 {
			secondSmallestNumber = n0
		}
	}

	if smallestNumber == secondSmallestNumber {
		return largestNumber - smallestNumber
	}

	if largestIndex == smallestIndex {
		return max(largestNumber-secondSmallestNumber, secondLargestNumber-smallestNumber)
	} else {
		return largestNumber - smallestNumber
	}
}
func main() {
	println(maxDistance([][]int{[]int{-3, -3}, []int{-3, -2}}))
}
