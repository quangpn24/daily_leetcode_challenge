package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type ListNode = utils.ListNode
type TreeNode = utils.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return &TreeNode{Val: head.Val}
    }
    slow, fast := head, head
    pre := head
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    root := &TreeNode{Val: slow.Val}
    pre.Next = nil
    root.Right = sortedListToBST(slow.Next)
    root.Left = sortedListToBST(head)
    return root
}
func main() {
    nodes := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{Val: 9}}}}}
    fmt.Println(sortedListToBST(nodes))
}
