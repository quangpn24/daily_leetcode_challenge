package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head
	tail := head
	length := 1
	//find tail
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	pre := head
	for i := 1; i <= length && head.Next != nil; i++ {
		if i%2 == 0 {
			//move to tail
			pre.Next = head.Next
			tail.Next = &ListNode{head.Val, nil}
			tail = tail.Next
		}
		pre = head
		head = head.Next
	}
	return res
}
func main() {
	head := &ListNode{1, &ListNode{1, nil}}
	res := oddEvenList(head)
	fmt.Println(res)
}
