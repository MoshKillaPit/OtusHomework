//	  Данные должны считываться в течение 1 минуты.
//	- Создайте горутину для обработки данных.
//
// Для каждых 10 полученных значений вычисляется среднее арифметическое и отправляется в канал с обработанными данными.
// - Главная горутина будет получать обработанные данные из канала и выводить их на экран.
// - Напишите юнит тесты на реализованные функции;
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func BufioScannerInput(numbersChan chan []int) {
	fmt.Println("Введите числа через пробел:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	b := scanner.Text()
	strValues := strings.Fields(b)
	if len(strValues) > 10 { // Ограничение до 10 чисел за один ввод
		strValues = strValues[:10]
	}
	size := make([]int, len(strValues))

	for i, str := range strValues {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка преобразования: " + err.Error())
			return
		}
		size[i] = num
	}
	timeout := time.After(1 * time.Minute)
	for {
		select {
		case numbersChan <- size: // Начинаем отправку

		case <-timeout: // Завершаем отправку спустя минуту
			return
		}
	}
}

func refactor(numbersChan chan []int) int {
	totalNumbers := <-numbersChan
	fmt.Println("Полученные данные", totalNumbers)
	totalCount := len(totalNumbers)
	sum := 0
	for _, num := range totalNumbers {
		sum += num
	}
	result := sum / totalCount
	return result
}

func main() {
	numbersChan := make(chan []int)
	go BufioScannerInput(numbersChan)
	go refactor(numbersChan)
	fmt.Println(refactor(numbersChan))
}
