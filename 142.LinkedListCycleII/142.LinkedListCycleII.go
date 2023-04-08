package main

import (
	"daily_leetcode_challenge/utils"
)

type ListNode = utils.ListNode

func detectCycle(head *ListNode) *ListNode {
	cycleMap := map[*ListNode]bool{}

	for head != nil {
		if _, ok := cycleMap[head]; ok {
			return head
		}
		cycleMap[head] = true
		head = head.Next
	}
	return nil
}
