package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        root = &TreeNode{Val: val}
        return root
    }
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    } else {
        root.Left = insertIntoBST(root.Left, val)
    }
    return root
}

func main() {
    root := &TreeNode{
        Val:   4,
        Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
        Right: &TreeNode{Val: 7},
    }
    fmt.Println(insertIntoBST(root, 5))
}
