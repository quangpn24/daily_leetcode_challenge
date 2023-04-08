package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
	"math"
)

type TreeNode = utils.TreeNode

func CheckBST(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return CheckBST(root.Right, root.Val, max) && CheckBST(root.Left, min, root.Val)

}
func isValidBST(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return CheckBST(root, math.MinInt, math.MaxInt)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}

	fmt.Println(isValidBST(root))
}
