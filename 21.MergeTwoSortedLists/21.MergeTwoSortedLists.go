package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type ListNode = utils.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resList := &ListNode{}
	tmpList := resList
	for list1 != nil || list2 != nil {
		if list1 == nil {
			tmpList.Next = list2
			break
		}
		if list2 == nil {
			tmpList.Next = list1
			break
		}
		if list1.Val > list2.Val {
			tmpList.Next = &ListNode{list2.Val, nil}
			list2 = list2.Next //increase step
			tmpList = tmpList.Next
		} else if list1.Val < list2.Val {
			tmpList.Next = &ListNode{list1.Val, nil}
			list1 = list1.Next // increase step
			tmpList = tmpList.Next
		} else {
			tmpList.Next = &ListNode{list1.Val, &ListNode{list2.Val, nil}}
			list1 = list1.Next // increase step
			list2 = list2.Next // increase step
			tmpList = tmpList.Next.Next
		}
	}
	return resList.Next
}
func main() {
	//list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	//list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list1 := &ListNode{1, nil}
	list2 := &ListNode{2, nil}
	resList := mergeTwoLists(list1, list2)
	fmt.Println(resList.Val)
}
