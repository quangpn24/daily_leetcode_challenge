package main

import "testing"

func Test_isValidBST(t *testing.T) {
	//root = [5,1,4,null,null,3,6]
	root1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	// root = [2,1,3]
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	// root = [5,4,6, null, null, 3, 7]
	root3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			"Test 1",
			root1,
			false,
		},
		{
			"Test 2",
			root2,
			true,
		},
		{
			"Test 3",
			root3,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
