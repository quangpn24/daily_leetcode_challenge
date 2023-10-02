package main

import (
    "daily_leetcode_challenge/utils"
)

type TreeNode = utils.TreeNode

func flatten(root *TreeNode) {
    if root != nil {
        flatten(root.Left)
        if root.Left != nil {
            tmp := root.Right
            root.Right = root.Left
            lastNode := FindLastRightNode(root.Left)
            lastNode.Right = tmp
            root.Left = nil
        }
        flatten(root.Right)
    }
}
func FindLastRightNode(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    for root.Right != nil {
        root = root.Right
    }
    return root
}

func main() {
    root := &TreeNode{
        Val:   1,
        Left:  &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}},
        Right: &TreeNode{5, &TreeNode{Val: 6}, &TreeNode{Val: 7}},
    }
    flatten(root)
}
