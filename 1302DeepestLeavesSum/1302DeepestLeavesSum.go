package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{root}
    sum := 0
    for len(queue) > 0 {
        sum = 0
        for _, node := range queue {
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            sum += node.Val
            queue = queue[1:]
        }
    }
    return sum
}
func main() {
    root := &TreeNode{
        Val:   1,
        Left:  &TreeNode{2, &TreeNode{4, &TreeNode{Val: 7}, nil}, &TreeNode{Val: 5}},
        Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 8}}},
    }
    fmt.Println(deepestLeavesSum(root))
}
