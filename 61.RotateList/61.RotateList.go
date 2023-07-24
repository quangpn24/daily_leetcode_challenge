package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	it := head
	copy := &ListNode{head.Val, nil}
	replica := copy
	length := 1
	for it.Next != nil {
		it = it.Next
		length++
		copy.Next = &ListNode{it.Val, nil}
		copy = copy.Next
	}
	copy.Next = head

	num := k % length
	for i := 0; i < length-num; i++ {
		replica = replica.Next
	}
	result := replica
	for i := 0; i < length-1; i++ {
		replica = replica.Next
	}
	replica.Next = nil
	return result
}
func main() {
	root := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	result := rotateRight(root, 4)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
