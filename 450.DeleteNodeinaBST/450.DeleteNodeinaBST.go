package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == key {
			return nil
		} else {
			return root
		}
	}
	res := root
	parentDeletedNode := root
	for root != nil {
		if root.Val == key {
			//delete node
			nodeInstead := root
			parentNodeInstead := parentDeletedNode
			if root.Right != nil {
				nodeInstead = root.Right
				for nodeInstead.Left != nil {
					parentNodeInstead = nodeInstead
					nodeInstead = nodeInstead.Left
				}
				if root.Right == nodeInstead {
					root.Right = nodeInstead.Right
				} else {
					parentNodeInstead.Left = nodeInstead.Right
				}
			} else if root.Left != nil {
				nodeInstead = root.Left
				for nodeInstead.Right != nil {
					parentNodeInstead = nodeInstead
					nodeInstead = nodeInstead.Right
				}
				if root.Left == nodeInstead {
					root.Left = nodeInstead.Left
				} else {
					parentNodeInstead.Right = nodeInstead.Left
				}
			}
			if root == nodeInstead && root.Val < parentDeletedNode.Val {
				parentDeletedNode.Left = nil
			} else if root == nodeInstead && root.Val > parentDeletedNode.Val {
				parentDeletedNode.Right = nil
			}
			root.Val = nodeInstead.Val
			return res
		} else if root.Val < key {
			parentDeletedNode = root
			root = root.Right
		} else {
			parentDeletedNode = root
			root = root.Left
		}
	}
	return res
}
func main() {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}},
		Right: &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}},
	}
	res := deleteNode(root, 5)
	fmt.Println(res)
}
