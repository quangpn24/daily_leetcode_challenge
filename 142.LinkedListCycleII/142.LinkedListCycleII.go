package main

import (
	utils "daily_leetcode_challenge"
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
