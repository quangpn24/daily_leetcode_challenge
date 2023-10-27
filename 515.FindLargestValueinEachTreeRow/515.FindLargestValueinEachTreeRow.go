package main

import (
	"daily_leetcode_challenge/utils"
	"math"
)

type TreeNode = utils.TreeNode

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0 {
		maxRow := math.MinInt
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if maxRow < node.Val {
				maxRow = node.Val
			}
			queue = queue[1:]
		}
		res = append(res, maxRow)
	}
	return res
}
func main() {

}
