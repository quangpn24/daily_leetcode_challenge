package main

import (
    "daily_leetcode_challenge/utils"
    "fmt"
)

type ListNode = utils.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    res := head
    var pre *ListNode
    for head != nil {
        next := head.Next
        for next != nil && next.Val == head.Val {
            next = next.Next
        }
        if head.Next != next && pre == nil {
            res = next
        } else if head.Next != next {
            pre.Next = next
        } else {
            pre = head
        }
        head = next
    }
    return res
}
func main() {
    head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
    fmt.Println(deleteDuplicates(head))
}
