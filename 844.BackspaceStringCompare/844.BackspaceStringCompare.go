package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	s = process(s)
	t = process(t)
	if len(s) != len(t) {
		return false
	}
	for i, _ := range s {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}
func process(s string) string {
	cnt := 0
	res := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			cnt++
		} else if cnt == 0 {
			res = append(res, s[i])
		} else {
			cnt--
		}
	}
	return string(res)
}
func main() {
	fmt.Println(process("ab#cd"))
}
