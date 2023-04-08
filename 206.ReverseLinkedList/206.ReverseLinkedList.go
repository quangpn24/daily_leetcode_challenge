package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

// solution 1
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

// solution2
// reverse(list) = first + reverser(remaining)
func reverseListByRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newHead := reverseListByRecursive(next)
	next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	res := reverseList(list)
	fmt.Println(res.Val)
}
