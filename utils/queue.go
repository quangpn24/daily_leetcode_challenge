package utils

type MyQueue[T comparable] struct {
	queue []T
}

func Constructor[T comparable]() MyQueue[T] {
	return MyQueue[T]{}
}

func (this *MyQueue[T]) Push(x T) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue[T]) Pop() T {
	peek := this.Peek()
	this.queue = this.queue[1:len(this.queue)]
	return peek
}

func (this *MyQueue[T]) Peek() T {
	return this.queue[0]
}

func (this *MyQueue[T]) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
