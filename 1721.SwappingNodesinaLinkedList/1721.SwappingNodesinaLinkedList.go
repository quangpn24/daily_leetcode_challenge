package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 1
	fast := head
	for fast != nil && fast.Next != nil {
		length += 2
		fast = fast.Next.Next
	}
	if fast == nil {
		length--
	}

	var start, end *ListNode
	cnt := 1
	curr := head
	for curr != nil {
		if cnt == k {
			start = curr
		}
		if cnt == length-k+1 {
			end = curr
		}
		curr = curr.Next
		cnt++
		if start != nil && end != nil {
			break
		}
	}

	//swap
	start.Val, end.Val = end.Val, start.Val

	return head
}
func main() {
	nodes := &ListNode{Val: 1, Next: &ListNode{2, nil}}
	tmp := swapNodes(nodes, 1)
	fmt.Println(tmp)
}
