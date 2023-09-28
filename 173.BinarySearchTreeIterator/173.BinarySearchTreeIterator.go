package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type TreeNode = utils.TreeNode
type BSTIterator struct {
    queue   []int
    currPos int
}

func Constructor(root *TreeNode) BSTIterator {
    if root == nil {
        return BSTIterator{}
    }
    res := BSTIterator{currPos: -1}
    res.queue = append(res.queue, Constructor(root.Left).queue...)
    res.queue = append(res.queue, root.Val)
    res.queue = append(res.queue, Constructor(root.Right).queue...)
    return res
}

func (this *BSTIterator) Next() int {
    this.currPos++
    return this.queue[this.currPos]
}

func (this *BSTIterator) HasNext() bool {
    return this.currPos < len(this.queue)-1
}
func main() {
    root := &TreeNode{
        Val:   7,
        Left:  &TreeNode{Val: 3},
        Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}},
    }
    it := Constructor(root)
    fmt.Println(it.Next())
    fmt.Println(it.Next())
    fmt.Println(it.HasNext())
    fmt.Println(it.Next())
    fmt.Println(it.HasNext())
    fmt.Println(it.Next())
    fmt.Println(it.HasNext())
    fmt.Println(it.Next())
    fmt.Println(it.HasNext())
}
