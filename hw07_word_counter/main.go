package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[^\w\sА-Яа-я]`) // Регистр допустимых символов

func punkDelete(text string) string {
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
	text := "How are you ? ! Are are : пирвет пирвет Пи Пи пи"
	fmt.Println(countWords(text))
}
