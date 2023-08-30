package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    res := [][]int{}
    for len(queue) > 0 {
        valueInLevel := []int{}
        for _, node := range queue {
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            valueInLevel = append(valueInLevel, node.Val)
            queue = queue[1:]
        }
        res = append(res, valueInLevel)
    }
    return res
}
func main() {
    root := &TreeNode{
        Val:   3,
        Left:  &TreeNode{Val: 9},
        Right: &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}},
    }
    fmt.Println(levelOrder(root))
}
