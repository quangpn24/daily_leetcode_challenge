package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
	"math"
)

type TreeNode = utils.TreeNode

func DFS(root *TreeNode, target int) []*TreeNode {
	if root == nil {
		return nil
	}
	res := []*TreeNode{}
	res = append(res, root)
	if root.Val > target {
		res = append(res, DFS(root.Left, target)...)
	} else if root.Val < target {
		res = append(res, DFS(root.Right, target)...)
	}
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestorOfP := DFS(root, p.Val)
	ancestorOfQ := DFS(root, q.Val)
	max := math.Min(float64(len(ancestorOfP)), float64(len(ancestorOfQ)))
	LCA := root
	for i := 0; i < int(max); i++ {
		if ancestorOfQ[i] == ancestorOfP[i] {
			LCA = ancestorOfP[i]
		} else {
			break
		}
	}
	return LCA
}
func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	p := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	q := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	tmp := lowestCommonAncestor(root, p, q)
	fmt.Println(tmp)
}
