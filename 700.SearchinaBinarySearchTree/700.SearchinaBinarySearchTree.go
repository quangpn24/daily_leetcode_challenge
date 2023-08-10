package main

import "daily_leetcode_challenge/utils"

type TreeNode = utils.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}

	}
	return nil
}
func main() {

}
