package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    res := [][]int{}
    for len(queue) > 0 {
        valueInLevel := []int{}
        for _, node := range queue {
            valueInLevel = append(valueInLevel, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            queue = queue[1:]
        }
        res = append(res, valueInLevel)
    }

    n := len(res)
    for i := 0; i < n/2; i++ {
        //swap
        tmp := res[n-i-1]
        res[n-1-i] = res[i]
        res[i] = tmp
    }
    return res
}
func main() {
    root := &TreeNode{
        Val:   3,
        Left:  &TreeNode{Val: 9},
        Right: &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}},
    }
    fmt.Println(levelOrderBottom(root))
}
