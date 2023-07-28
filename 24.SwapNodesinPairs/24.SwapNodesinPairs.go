package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	var currNode *ListNode
	if head.Next != nil {
		newHead = head.Next
		head.Next = head.Next.Next
		newHead.Next = head
		currNode = head
		head = head.Next
	}
	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		currNode.Next = next
		currNode = head
		head = head.Next
	}
	return newHead
}
func main() {
	root := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	fmt.Println(swapPairs(root))
}
