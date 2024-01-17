package main

import (
	"fmt"
	"strings"
)

func removeOccurrences(s string, part string) string {
	pos := strings.Index(s, part)
	for s != "" && pos != -1 {
		s = s[:pos] + s[pos+len(part):]
		pos = strings.Index(s, part)
	}
	return s
}
func main() {
	fmt.Println(removeOccurrences("eemckxmckx", "emckx"))
}
