package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxL := 1
	maxR := 1
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		maxL = maxL + maxDepth(root.Left)
	}
	if root.Right != nil {
		maxR = maxR + maxDepth(root.Right)
	}
	if maxL > maxR {
		return maxL
	} else {
		return maxR
	}
}
func main() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}},
		Right: &TreeNode{1, &TreeNode{4, nil, nil}, &TreeNode{2, &TreeNode{9, nil, nil}, &TreeNode{11, &TreeNode{8, nil, nil}, &TreeNode{10, nil, nil}}}},
	}
	fmt.Println(maxDepth(root1))
}
