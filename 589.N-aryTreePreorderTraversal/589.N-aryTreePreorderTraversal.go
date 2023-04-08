package main

import (
	"daily_leetcode_challenge/utils"
)

type Node = utils.Node

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	for _, c := range root.Children {
		tmp := preorder(c)
		res = append(res, tmp...)
	}
	return res
}
