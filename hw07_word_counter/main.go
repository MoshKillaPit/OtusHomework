//- Напишите функцию countWords, которая принимает на вход строку текста и возвращает мапу, содержащую количество упоминаний каждого слова в тексте.
//- Вы можете использовать мапу для отображения слова в его частоту появления в тексте
//- Вы также можете использовать слайс и функции для работы со строками, чтобы разделить текст на отдельные слова и очистить
//  их от пунктуации и пробелов перед подсчетом;
//- Напишите юнит тест на реализованную функцию;

package main

import (
	"fmt"
	"strings"
)

var text string

//var cache map[string]int

// func countWords (string) map[string]int {}
// Функция разбития строки на массив строк
func stringGet(string) {
	fmt.Println("Пихни суда текст")
	fmt.Scanf("%s", &text)
	Words := strings.Fields(text)
	fmt.Println(Words)
	return
}
func main() {
	massivString := stringGet
	fmt.Println(massivString)
}
