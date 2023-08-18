package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	maxSumInLevel := root.Val
	level := 1
	res := 1
	for len(queue) > 0 {
		sum := 0
		for _, node := range queue {
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		if sum > maxSumInLevel {
			maxSumInLevel = sum
			res = level
		}
		level++
	}
	return res
}
func main() {
	root := &TreeNode{
		Val:   -100,
		Left:  &TreeNode{-200, &TreeNode{Val: -20}, &TreeNode{Val: -5}},
		Right: &TreeNode{-300, &TreeNode{Val: -10}, nil},
	}
	fmt.Println(maxLevelSum(root))
}
