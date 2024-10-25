package main

import "daily_leetcode_challenge/utils"

type TreeNode = utils.TreeNode

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	var (
		isFlip = true
		noFlip = true
	)

	noFlip = flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)

	isFlip = flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)

	return noFlip || isFlip
}

func main() {
	print(min(1, 0, -1))
}
