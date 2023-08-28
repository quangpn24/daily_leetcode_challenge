package main

import "fmt"

func removeDuplicates(nums []int) int {
    myMap := make(map[int]bool)

    index := 0
    for i := 0; i < len(nums); i++ {
        if _, ok := myMap[nums[i]]; ok {
            continue
        } else {
            myMap[nums[i]] = true
            nums[index] = nums[i]
            index++
        }
    }
    return len(myMap)
}
func main() {
    nums := []int{1, 1, 2}
    fmt.Println(removeDuplicates(nums))
}
