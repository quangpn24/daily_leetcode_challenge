package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func sumNumbers(root *TreeNode) int {
    sum := 0
    if root.Left != nil {
        sum += DFS(root.Left, root.Val)
    }
    if root.Right != nil {
        sum += DFS(root.Right, root.Val)
    }
    if sum == 0 {
        return root.Val
    }
    return sum
}
func DFS(root *TreeNode, sum int) int {
    sum1, sum2 := 0, 0
    tmp := sum*10 + root.Val
    if root.Left != nil {
        sum1 = DFS(root.Left, tmp)
    }
    if root.Right != nil {
        sum2 = DFS(root.Right, tmp)
    }
    if sum1 == sum2 && sum1 == 0 {
        return tmp
    }
    return sum1 + sum2
}
func main() {
    root := &TreeNode{
        Val:   4,
        Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
        Right: &TreeNode{Val: 0},
    }
    fmt.Println(sumNumbers(root))
}
