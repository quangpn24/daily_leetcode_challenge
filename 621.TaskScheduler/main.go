package main

import "container/heap"

type Item struct {
	Value    string
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest priority, so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the heap.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority item.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (piq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(piq, item.Index)
}
func leastInterval(tasks []string, n int) int {
	queue := make(PriorityQueue, 0)

	// Push items into the priority queue.
	heap.Init(&queue)

	m := make(map[string]int)
	for _, task := range tasks {
		m[task]++
	}

	for key, value := range m {
		heap.Push(&queue, &Item{
			Value:    string(key),
			Priority: value,
		})
	}

	cnt := 0
	for queue.Len() > 0 {
		circle := make(map[string]int)
		for i := 0; i <= n; i++ {
			if queue.Len() <= 0 {
				println("idle"+"   ", cnt)
				break
			}

			item := heap.Pop(&queue).(*Item)
			println(item.Value, "   ", cnt)
			circle[item.Value] = item.Priority - 1
		}
		cnt += n + 1

		//push to queue
		for k, v := range circle {
			if v == 0 {
				continue
			}
			heap.Push(&queue, &Item{
				Value:    k,
				Priority: v,
			})
		}
		if queue.Len() == 0 {
			cnt -= n + 1 - len(circle)
		}
	}

	return cnt
}
func main() {
	//println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
	//println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	println(leastInterval([]string{"A", "A", "B", "B", "C", "C", "D", "D", "E", "E"}, 4))
	//println(leastInterval([]string{"A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G"}, 1))
	//println(leastInterval([]string{"F", "J", "J", "A", "J", "F", "C", "H", "J", "B", "E", "G", "G", "F", "A", "C", "I", "F", "J", "C", "J", "C", "H", "C", "A", "D", "G", "H", "B", "F", "G", "C", "C", "A", "E", "B", "H", "J", "E", "I", "F", "D", "E", "A", "C", "D", "B", "D", "J", "J", "C", "F", "D", "D", "J", "H", "A", "E", "C", "D", "J", "D", "G", "G", "B", "C", "E", "G", "H", "I", "D", "H", "F", "E", "I", "B", "D", "E", "I", "E", "C", "J", "G", "I", "D", "E", "D", "J", "C", "A", "C", "C", "D", "I", "J", "B", "D", "H", "H", "J", "G", "B", "G", "A", "H", "E", "H", "E", "D", "E", "J", "E", "J", "C", "F", "C", "J", "G", "B", "C", "I", "I", "H", "F", "A", "D", "G", "F", "C", "C", "F", "G", "C", "J", "B", "B", "I", "C", "J", "J", "E", "G", "H", "C", "I", "G", "J", "I", "G", "G", "J", "G", "G", "E", "G", "B", "I", "J", "B", "H", "D", "H", "G", "F", "C", "H", "C", "D", "A", "G", "B", "H", "H", "B", "J", "C", "A", "F", "J", "G", "F", "E", "B", "F", "E", "B", "B", "A", "E", "F", "E", "H", "I", "I", "C", "G", "J", "D", "H", "E", "F", "G", "G", "D", "E", "B", "F", "J", "J", "J", "D", "H", "E", "B", "D", "J", "I", "F", "C", "I", "E", "H", "F", "E", "G", "D", "E", "C", "F", "E", "D", "E", "A", "I", "E", "A", "D", "H", "G", "C", "I", "E", "G", "A", "H", "I", "G", "G", "A", "G", "F", "H", "J", "D", "F", "A", "G", "H", "B", "J", "A", "H", "B", "H", "C", "G", "F", "A", "C", "C", "B", "I", "G", "G", "B", "C", "J", "J", "I", "E", "G", "D", "I", "J", "I", "C", "G", "A", "J", "G", "F", "J", "F", "C", "F", "G", "J", "I", "E", "B", "G", "F", "A", "D", "A", "I", "A", "E", "H", "F", "D", "D", "C", "B", "J", "I", "J", "H", "I", "C", "D", "A", "G", "F", "I", "B", "E", "D", "C", "J", "G", "I", "H", "E", "C", "E", "I", "I", "B", "B", "H", "J", "C", "F", "I", "D", "B", "F", "H", "F", "A", "C", "A", "A", "B", "D", "C", "A", "G", "B", "G", "F", "E", "G", "A", "A", "A", "C", "J", "H", "H", "G", "C", "C", "B", "C", "E", "B", "E", "F", "I", "E", "E", "D", "I", "H", "G", "F", "A", "H", "B", "J", "B", "G", "H", "C", "C", "B", "G", "C", "B", "A", "E", "G", "A", "J", "G", "D", "C", "I", "G", "F", "G", "G", "A", "J", "E", "I", "D", "E", "A", "F", "A", "H", "C", "E", "D", "D", "D", "H", "I", "F", "F", "A", "F", "A", "A", "C", "J", "D", "J", "H", "I", "F", "A", "C", "B", "C", "A", "C", "C", "H", "A", "J", "I", "B", "A", "I", "F", "J", "C", "I", "B", "C", "E", "E", "E", "J", "G", "F", "E", "I", "A", "A", "E", "B", "J", "H", "H", "H", "A", "H", "J", "E", "F", "E", "F", "G", "J", "D", "I", "D", "I", "F", "B", "J", "D", "A", "A", "D", "F", "G", "B", "J", "H", "F", "A", "D", "H", "C", "B", "A", "J", "H", "I", "F", "H", "E", "G", "B", "A", "F", "F", "A", "C", "D", "G", "I", "I", "J", "H", "H", "C", "J", "G", "B", "A", "D", "B", "F", "J", "D", "I", "A", "F", "F", "F", "F", "A", "E", "B", "C", "G", "H", "E", "B", "B", "A", "G", "D", "C", "C", "E", "A", "C", "F", "G", "A", "I", "F", "B", "H", "J", "G", "C", "B", "H", "D", "A", "H", "B", "H", "H", "C", "A", "F", "I", "C", "F", "A", "C", "J", "I", "H", "H", "F", "B", "B", "D", "E", "C", "J", "F", "C", "E", "A", "J", "E", "C", "A", "E", "B", "A", "J", "F", "J", "J", "J", "H", "H", "C", "I", "E", "G", "G", "H", "J", "J", "H", "H", "H", "J", "H", "A", "G", "I", "C", "E", "C", "D", "G", "G", "F", "H", "D", "G", "H", "A", "E", "I", "D", "A", "H", "G", "E", "A", "B", "F", "I", "C", "A", "F", "B", "A", "I", "F", "G", "I", "F", "D", "A", "B", "J", "B", "D", "F", "G", "J", "J", "A", "A", "C", "H", "G", "F", "B", "I", "I", "J", "A", "H", "D", "F", "E", "F", "J", "B", "F", "C", "G", "E", "A", "G", "H", "E", "H", "H", "F", "I", "G", "C", "C", "G", "J", "B", "H", "F", "H", "D", "I", "B", "D", "I", "F", "H", "I", "D", "F", "G", "G", "E", "A", "C", "A", "G", "H", "G", "H", "J", "F", "D", "F", "G", "D", "D", "C", "J", "C", "J", "G", "G", "G", "G", "H", "H", "G", "D", "E", "H", "G", "C", "B", "F", "I", "F", "C", "H", "J", "I", "A", "F", "D", "C", "F", "C", "E", "E", "D", "D", "C", "G", "B", "F", "E", "J", "C", "I", "E", "D", "B", "B", "I", "I", "I", "H", "C", "E", "C", "J", "F", "G", "A", "I", "J", "D", "I", "C", "G", "F", "I", "E", "I", "E", "F", "A", "G", "E", "J", "A", "I", "A", "D", "A", "G", "J", "F", "E", "D", "I", "A", "E", "J", "I", "C", "J", "B", "F", "B", "E", "C", "E", "F", "G", "E", "J", "J", "I", "E", "D", "F", "C", "H", "H", "B", "G", "D", "I", "I", "F", "B", "G", "C", "F", "J", "B", "G", "J", "H", "D", "G", "C", "C", "I", "I", "E", "I", "B", "H", "B", "I", "G", "F", "H", "G", "C", "J", "D", "C", "E", "G", "F", "C", "H", "D", "A", "C", "D", "H", "B", "C", "H", "I", "B", "A", "J", "C", "B", "D", "J", "D", "H", "F", "B", "A", "G", "G", "J", "I", "E", "F", "A", "D", "H", "D", "B", "C", "A", "H", "F", "G", "B", "F", "H", "B", "H", "I", "J", "D", "H", "I", "B", "C", "D", "G", "A", "E", "A", "A", "I", "F", "I", "F", "B", "B", "I", "F", "A", "E", "I", "A", "B", "G", "C", "F", "I", "A", "F", "I", "D", "H", "B", "I", "I", "B", "J", "F", "E", "B", "B", "B", "D", "C", "J", "E", "J", "J", "G", "D", "F", "F", "F", "G", "I", "H", "J", "J", "G", "D", "G", "F"}, 8))
}
