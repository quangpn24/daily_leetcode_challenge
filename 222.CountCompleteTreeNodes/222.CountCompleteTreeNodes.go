package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    count := 0
    count += countNodes(root.Left) + countNodes(root.Right) + 1
    return count
}
func main() {
    root := &TreeNode{
        Val:   1,
        Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
        Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: nil},
    }
    fmt.Println(countNodes(root))
}
