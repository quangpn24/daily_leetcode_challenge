package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
	"math"
)

type TreeNode = utils.TreeNode

func findAncestor(root *TreeNode, target int) []*TreeNode {
	res := []*TreeNode{}
	if root == nil {
		return nil
	}
	if root.Val == target {
		return append(res, root)
	}
	if root.Left != nil || root.Right != nil {
		res = append(res, root)
	}
	if root.Left != nil {
		resl := findAncestor(root.Left, target)
		if len(resl) > 0 && resl[len(resl)-1].Val == target {
			return append(res, resl...)
		}
	}
	if root.Right != nil {
		resr := findAncestor(root.Right, target)
		if len(resr) > 0 && resr[len(resr)-1].Val == target {
			return append(res, resr...)
		}
	}
	return res
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	arrP := findAncestor(root, p.Val)
	arrQ := findAncestor(root, q.Val)
	minLen := math.Min(float64(len(arrP)), float64(len(arrQ)))
	i := 0
	for ; i < int(minLen); i++ {
		if arrQ[i] != arrP[i] {
			break
		}
	}
	return arrP[i-1]
}
func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	p := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	q := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	tmp := lowestCommonAncestor(root, p, q)
	fmt.Println(tmp)
}
