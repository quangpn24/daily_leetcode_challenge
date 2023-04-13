package main

import (
	"daily_leetcode_challenge/utils"
	"math"
)

type TreeNode = utils.TreeNode

func BFS(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	depthLeft, check1 := BFS(root.Left)
	depthRight, check2 := BFS(root.Right)
	depthLeft++
	depthRight++
	if !check1 || !check2 {
		return 0, false
	}
	if math.Abs(float64(depthLeft-depthRight)) > 1 {
		return 0, false
	}
	return int(math.Max(float64(depthLeft), float64(depthRight))), true
}

func isBalanced(root *TreeNode) bool {
	_, ok := BFS(root)
	return ok
}
