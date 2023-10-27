package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	pre, curr, next := head, head, head.Next
	for next != nil {
		if curr.Val == val {
			pre.Next = next
		} else {
			pre = curr
		}
		curr = next
		next = curr.Next
	}
	if curr.Val == val {
		pre.Next = next
	}
	return head
}
func main() {
	head := &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}
	tmp := removeElements(head, 7)
	fmt.Println(tmp)
}
