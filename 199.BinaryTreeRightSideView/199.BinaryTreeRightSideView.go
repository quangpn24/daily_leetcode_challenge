package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := []int{}
	i := 1
	length := 1
	for ; i <= length; i++ {
		pick := queue[0]
		queue = queue[1:]
		if pick.Left != nil {
			queue = append(queue, pick.Left)
		}
		if pick.Right != nil {
			queue = append(queue, pick.Right)
		}
		if i == length {
			res = append(res, pick.Val)
			length += len(queue)
		}
	}
	return res
}
func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}},
		Right: &TreeNode{3, nil, &TreeNode{Val: 4}},
	}
	fmt.Println(rightSideView(root))
}
