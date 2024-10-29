package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	GroupNumber int
	Average     int
}

func randomNum(randomNumbers chan int, limit int, duration time.Duration) {
	defer close(randomNumbers)
	timeout := time.After(duration)

	for count := 0; count < limit; count++ {
		select {
		case <-timeout:
			fmt.Println("Время приёма данных истекло")
			return
		default:
			numRan := rand.Intn(100) // #nosec G404
			randomNumbers <- numRan
			time.Sleep(1 * time.Second)
		}
	}
}

func refactor(randomNumbers chan int, resultChan chan Result) {
	defer close(resultChan)

	numMas := []int{}
	groupNumber := 1

	for num := range randomNumbers {
		numMas = append(numMas, num)

		if len(numMas) == 10 {
			sum := 0
			for _, n := range numMas {
				sum += n
			}
			average := sum / len(numMas)
			resultChan <- Result{GroupNumber: groupNumber, Average: average}

			numMas = []int{}
			groupNumber++
		}
	}

	// Отправляем остаток, если чисел меньше 10
	if len(numMas) > 0 {
		sum := 0
		for _, n := range numMas {
			sum += n
		}
		average := sum / len(numMas)
		resultChan <- Result{GroupNumber: groupNumber, Average: average}
	}
}

func main() {
	randomNumbers := make(chan int)
	resultChan := make(chan Result)
	duration := 1 * time.Minute

	// Запуск горутины для генерации случайных чисел
	go randomNum(randomNumbers, 100, duration)

	// Запуск горутины для обработки данных
	go refactor(randomNumbers, resultChan)

	// Получение и вывод результата
	for result := range resultChan {
		fmt.Printf("Среднее арифметическое группы #%d: %d\n", result.GroupNumber, result.Average)
	}
}
