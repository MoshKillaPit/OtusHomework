package main

import (
	"fmt"
	"regexp"
	"strings"
)

func punkDelete(text string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	stringClear := re.ReplaceAllString(text, "") // Удалили пунктацию
	return stringClear
}

func lowClear(stringClear string) []string {
	lowClear := strings.ToLower(stringClear)
	words := strings.Fields(lowClear)
	return words
}

func countWords(text string) map[string]int {
	text1 := punkDelete(text)
	text2 := lowClear(text1)
	m := make(map[string]int)
	for _, word := range text2 {
		m[word]++
	}
	return m
}

func main() {
	text := "How are you ? ! Are are :"
	fmt.Println(countWords(text))
}
