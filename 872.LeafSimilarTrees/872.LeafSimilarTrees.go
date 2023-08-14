package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func NodeLeaf(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	if root.Left != nil {
		res = append(res, NodeLeaf(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, NodeLeaf(root.Right)...)
	}
	return res
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1 := NodeLeaf(root1)
	res2 := NodeLeaf(root2)
	if len(res1) != len(res2) {
		return false
	}
	for i := 0; i < len(res1); i++ {
		if res1[i] != res2[i] {
			return false
		}
	}
	return true
}
func main() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}},
		Right: &TreeNode{1, &TreeNode{4, nil, nil}, &TreeNode{2, &TreeNode{9, nil, nil}, &TreeNode{11, &TreeNode{8, nil, nil}, &TreeNode{10, nil, nil}}}},
	}
	root2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}},
		Right: &TreeNode{1, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}},
	}
	fmt.Println(leafSimilar(root1, root2))
}
