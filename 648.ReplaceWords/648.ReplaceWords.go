package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	arr := strings.Split(sentence, " ")
	for i, val := range arr {
		var tmp string
		for _, d := range dictionary {
			if strings.HasPrefix(val, d) {
				if len(tmp) == 0 {
					tmp = d
				} else if len(d) < len(tmp) {
					tmp = d
				}
			}
		}
		if len(tmp) > 0 {
			arr[i] = tmp
		}
	}
	return strings.Join(arr, " ")
}
func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
