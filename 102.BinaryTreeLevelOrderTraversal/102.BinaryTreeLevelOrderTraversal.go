package main

import (
	"daily_leetcode_challenge/utils"
	"fmt"
)

type TreeNode = utils.TreeNode

func enqueue(queue []*TreeNode, element *TreeNode) []*TreeNode {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		return element, []*TreeNode{}

	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}

func levelOrder(root *TreeNode) [][]int {

}
