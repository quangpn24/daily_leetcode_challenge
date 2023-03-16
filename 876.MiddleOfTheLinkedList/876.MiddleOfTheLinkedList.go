package main

import (
	utils "daily_leetcode_challenge"
	"fmt"
)

type ListNode = utils.ListNode

func middleNode(head *ListNode) *ListNode {
	pointSlow := head
	pointerFast := head // step x2

	for pointerFast != nil && pointerFast.Next != nil {
		pointSlow = pointSlow.Next
		pointerFast = pointerFast.Next.Next
	}

	return pointSlow
}

func main() {
	//list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	resList := middleNode(list)
	fmt.Println(resList.Val)
}
