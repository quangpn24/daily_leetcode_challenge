package main

import "fmt"

type Node struct {
    Val   int
    Left  *Node
    Right *Node
    Next  *Node
}

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
            if i == n-1 {
                queue[i].Next = nil
            } else {
                queue[i].Next = queue[i+1]
            }
        }
        queue = queue[n:]
    }
    return root
}
func main() {
    root := &Node{
        Val:   1,
        Left:  &Node{2, &Node{Val: 4}, &Node{Val: 5}, nil},
        Right: &Node{3, &Node{Val: 6}, &Node{Val: 7}, nil},
    }
    fmt.Println(connect(root))
}
