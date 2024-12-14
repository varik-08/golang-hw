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

	words := strings.Split(cleanStr, " ")

	for _, word := range words {
		if count := mapWords[word]; count > 0 {
			mapWords[word]++
		} else {
			mapWords[word] = 1
		}
	}

	return mapWords
}