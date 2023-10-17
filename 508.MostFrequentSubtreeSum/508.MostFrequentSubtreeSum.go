package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
    "math"
)

type TreeNode = utils.TreeNode

func findFrequentTreeSum(root *TreeNode) []int {
    m := make(map[int]int)
    s := dfs(root, m)
    m[s]++

    valMax := math.MinInt
    res := make([]int, 0)
    for k, v := range m {
        if v > valMax {
            valMax = v
            res = []int{k}
            continue
        }
        if v == valMax {
            res = append(res, k)
        }
    }
    return res
}
func dfs(root *TreeNode, m map[int]int) int {
    res := 0
    if root != nil {
        res += root.Val
    }
    if root.Left != nil {
        valueL := dfs(root.Left, m)
        m[valueL]++
        res += valueL
    }
    if root.Right != nil {
        valueR := dfs(root.Right, m)
        m[valueR]++
        res += valueR
    }
    return res
}
func main() {
    root := &TreeNode{
        Val:   5,
        Left:  &TreeNode{Val: 2},
        Right: &TreeNode{Val: -5},
    }
    fmt.Println(findFrequentTreeSum(root))
}
