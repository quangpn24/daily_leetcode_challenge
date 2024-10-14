package main

import (
	"strconv"
	"strings"

	"daily_leetcode_challenge/utils"
)

type TreeNode = utils.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	return try(root, []string{})
}

func try(root *TreeNode, path []string) []string {
	if root == nil {
		return []string{}
	}
	path = append(path, strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		return []string{strings.Join(path, "->")}
	}
	result := []string{}
	result = append(result, try(root.Left, path)...)
	result = append(result, try(root.Right, path)...)

	path = path[:len(path)-1]
	return result
}

func main() {
	res := binaryTreePaths(&TreeNode{1, &TreeNode{2, nil, &TreeNode{Val: 5}}, &TreeNode{Val: 3}})
	for _, re := range res {
		println(re)
	}
}
