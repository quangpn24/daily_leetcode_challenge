package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return nil
    }
    res = append(res, root.Val)
    res = append(res, preorderTraversal(root.Left)...)
    res = append(res, preorderTraversal(root.Right)...)
    return res
}

func main() {
    root := &TreeNode{Val: 1,
    }
    fmt.Println(preorderTraversal(root))
}
