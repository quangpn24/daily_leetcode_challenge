package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
    arr := strings.Split(sentence, " ")
    for index, word := range arr {
        if len(word) >= len(searchWord) && word[0:len(searchWord)] == searchWord {
            return index + 1
        }
    }
    return -1
}
func main() {

}
