package utils

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}
func (s *Stack[T]) Pop() T {
	var emptyT T
	if s.IsEmpty() {
		return emptyT
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
func (s *Stack[T]) Top() T {
	var emptyT T
	if s.IsEmpty() {
		return emptyT
	}
	return (*s)[len(*s)-1]
}
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
