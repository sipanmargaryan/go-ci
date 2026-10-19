package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	counts := make(map[string]int)

	for _, word := range strings.Fields(strings.ToLower(text)){
		counts[word]++
	}

	return  counts
}	

func main() {
	fmt.Println(countWords("Go is good Go"))
	fmt.Println(countWords(""))
}