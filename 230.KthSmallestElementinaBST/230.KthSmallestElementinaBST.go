package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode

func kthSmallest(root *TreeNode, k int) int {
    res := []int{}
    dfs(root, k, &res)
    return res[len(res)-1]
}
func dfs(root *TreeNode, k int, res *[]int) {
    if root == nil || len(*res) == k {
        return
    }
    dfs(root.Left, k, res)
    if len(*res) < k {
        *res = append(*res, root.Val)
    }
    dfs(root.Right, k, res)
}
func main() {
    root := &TreeNode{
        Val:   5,
        Left:  &TreeNode{3, &TreeNode{2, &TreeNode{Val: 1}, nil}, &TreeNode{Val: 4}},
        Right: &TreeNode{Val: 6},
    }
    fmt.Println(kthSmallest(root, 1))
}
