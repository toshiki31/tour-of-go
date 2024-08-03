package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// string s で渡される文章の、各単語の出現回数のmapを返す

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for _, word := range words {
		_, ok := wordCount[word]
		if ok {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
