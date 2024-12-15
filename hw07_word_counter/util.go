package main

import (
	"regexp"
	"strings"
)

func removePunctuation(input string) string {
	re := regexp.MustCompile(`[[:punct:]]`)
	return re.ReplaceAllString(input, "")
}

func CountWords(str string) map[string]int {
	mapWords := make(map[string]int)
	cleanStr := removePunctuation(str)

	words := strings.Fields(cleanStr)

	for _, word := range words {
		mapWords[word]++
	}

	return mapWords
}
