package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]}
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    return root
}
func main() {
    fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
