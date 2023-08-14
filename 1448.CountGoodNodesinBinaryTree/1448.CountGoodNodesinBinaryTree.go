package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func findGoodNode(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	numOfNodes := 0
	if max <= root.Val {
		max = root.Val
		numOfNodes++
	}
	if root.Left != nil {
		numOfNodes += findGoodNode(root.Left, max)
	}
	if root.Right != nil {
		numOfNodes += findGoodNode(root.Right, max)
	}
	return numOfNodes
}
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := root.Val
	numOfNode := 1
	if root.Left != nil {
		numOfNode += findGoodNode(root.Left, max)
	}
	if root.Right != nil {
		numOfNode += findGoodNode(root.Right, max)
	}
	return numOfNode
}
func main() {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{1, &TreeNode{3, nil, nil}, nil},
		Right: &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{5, nil, nil}},
	}
	fmt.Println(goodNodes(root))
}
