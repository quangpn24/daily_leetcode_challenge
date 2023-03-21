package main

import utils "daily_leetcode_challenge"

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
