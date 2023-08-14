package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func pairSum(head *ListNode) int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val, head.Next.Val)
		head = head.Next.Next
	}
	max := 0
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i]+arr[n-1-i] > max {
			max = arr[i] + arr[n-1-i]
		}
	}
	return max
}
func main() {
	head := &ListNode{1, &ListNode{100001, nil}}
	fmt.Println(pairSum(head))
}
