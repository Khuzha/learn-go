package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		_, ok := result[word]
		if (!ok) {
			result[word] = 1
		} else {
			result[word]++
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
