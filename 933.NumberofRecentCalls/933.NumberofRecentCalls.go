package main

import "fmt"

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	for len(this.queue) > 0 {
		if t-3000 <= this.queue[0] {
			return len(this.queue)
		} else {
			this.queue = this.queue[1:len(this.queue)]
		}
	}
	return 0
}

func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(642))
	fmt.Println(rc.Ping(1849))
	fmt.Println(rc.Ping(4921))
	fmt.Println(rc.Ping(5936))
	fmt.Println(rc.Ping(5957))

}
