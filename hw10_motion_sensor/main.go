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
	defer close(numbersChan) // Закрываем канал после завершения функции
	timeout := time.After(1 * time.Minute)
	fmt.Println("Введите числа через пробел:")

	for {
		select {
		case <-timeout: // Завершаем ввод спустя минуту
			fmt.Println("Время для ввода истекло")
			return
		default:
			// Считывание данных
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				b := scanner.Text()
				strValues := strings.Fields(b)
				if len(strValues) > 10 { // Ограничение до 10 чисел за один ввод
					strValues = strValues[:10]
				}
				size := make([]int, 0, len(strValues)) // Инициализируем срез
				for _, str := range strValues {
					num, err := strconv.Atoi(str)
					if err != nil {
						fmt.Println("Ошибка преобразования: " + err.Error())
						continue // Пропускаем некорректные значения
					}
					size = append(size, num) // Добавляем только корректные значения
				}
				numbersChan <- size
			}
		}
	}
}

func refactor(numbersChan chan []int, resultChan chan int) {
	defer close(resultChan) // Закрываем канал после завершения работы горутины
	count := 0
	sum := 0
	for totalNumbers := range numbersChan {
		count += len(totalNumbers)
		for _, num := range totalNumbers {
			sum += num
		}
		// Отправляем результат каждые 10 чисел
		if count >= 10 {
			resultChan <- sum / count
			count = 0
			sum = 0
		}
	}
	if count > 0 {
		resultChan <- sum / count
	}
}

func main() {
	numbersChan := make(chan []int)
	resultChan := make(chan int)

	go BufioScannerInput(numbersChan)    // Запустили горутину по сбору чисел
	go refactor(numbersChan, resultChan) // Запустили обработку

	for result := range resultChan {
		fmt.Println("Результат:", result)
	}

	fmt.Println("Программа завершена")
}
