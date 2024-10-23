//	  Данные должны считываться в течение 1 минуты.
//	- Создайте горутину для обработки данных.
//
// Для каждых 10 полученных значений вычисляется среднее арифметическое и отправляется в канал с обработанными данными.
// - Главная горутина будет получать обработанные данные из канала и выводить их на экран.
// - Напишите юнит тесты на реализованные функции;
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	SliceNumber int
	Average     int
}

type SliceData struct {
	SliceNumber int
	Numbers     []int
}

func randomNum(randomNumbers chan SliceData, limit int) {
	defer close(randomNumbers)
	timeout := time.After(1 * time.Minute)

	numMas := []int{} // Один срез, в который пишем данные
	sliceNumber := 1  // Переменная для нумерации срезов

	for count := 0; count < limit; count++ {
		select {
		case <-timeout:
			fmt.Println("Время приёма данных истекло")
			if len(numMas) > 0 { // Добавление оставшихся чисел после таймера
				randomNumbers <- SliceData{SliceNumber: sliceNumber, Numbers: numMas}
			}
			return
		default:
			numRan := rand.Intn(100)
			numMas = append(numMas, numRan) // Запись чисел в срез

			if len(numMas) == 10 {
				// Отправляем срез вместе с его номером в канал
				randomNumbers <- SliceData{SliceNumber: sliceNumber, Numbers: numMas}
				fmt.Printf("Отправлен срез #%d: %v\n", sliceNumber, numMas)
				numMas = []int{} // Обнуляем срез для новых чисел
				sliceNumber++    // Увеличиваем номер среза
			}
		}
	}
	// Отправляем остаток данных
	if len(numMas) > 0 {
		randomNumbers <- SliceData{SliceNumber: sliceNumber, Numbers: numMas}
		fmt.Printf("Отправлен срез #%d: %v\n", sliceNumber, numMas)
	}
}

func refactor(randomNumbers chan SliceData, resultChan chan Result) {
	defer close(resultChan)

	// Получаем данные из канала
	for sliceData := range randomNumbers {
		count := len(sliceData.Numbers)
		sum := 0

		// Считаем сумму элементов среза
		for _, num := range sliceData.Numbers {
			sum += num
		}

		// Отправляем результат обработки с номером среза
		if count > 0 {
			resultChan <- Result{SliceNumber: sliceData.SliceNumber, Average: sum / count}
		} else {
			resultChan <- Result{SliceNumber: sliceData.SliceNumber, Average: 0}
		}
	}
}

func main() {
	randomNumbers := make(chan SliceData)
	resultChan := make(chan Result)

	// Запуск горутины для генерации случайных чисел
	go randomNum(randomNumbers, 100)

	// Запуск горутины для обработки данных
	go refactor(randomNumbers, resultChan)

	// Получение и вывод результата
	for result := range resultChan {
		fmt.Printf("Среднее арифметическое среза #%d: %d\n", result.SliceNumber, result.Average)
	}
}
