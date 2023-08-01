package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	curr := head
	doubleJump := head
	pre := head
	for doubleJump.Next != nil && doubleJump.Next.Next != nil {
		pre = curr
		curr = curr.Next
		doubleJump = doubleJump.Next.Next
	}
	if doubleJump.Next != nil {
		pre = curr
		curr = curr.Next
	}

	//delete
	pre.Next = curr.Next
	return head
}
func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	res := deleteMiddle(head)
	fmt.Println(res.Val)
}
