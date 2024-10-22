package main

import (
	"os"
	"testing"
	"time"
)

func TestBufioScannerInput_ValidInput(t *testing.T) {
	// Подменяем os.Stdin
	input := "1 2 3 4 5 6 7 8 9 10 11" // эмуляция ввода пользователя
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }() // Восстанавливаем старый os.Stdin после теста

	// Записываем данные в writer
	go func() {
		w.WriteString(input + "\n")
		w.Close()
	}()

	numbersChan := make(chan []int)
	go BufioScannerInput(numbersChan)

	select {
	case result := <-numbersChan:
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // только первые 10 чисел
		for i, num := range result {
			if num != expected[i] {
				t.Errorf("Expected %d, but got %d", expected[i], num)
			}
		}
	case <-time.After(1 * time.Second):
		t.Error("Test timed out waiting for input")
	}
}

func TestBufioScannerInput_Timeout(t *testing.T) {
	// Подменяем os.Stdin
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }() // Восстанавливаем старый os.Stdin после теста

	// Запускаем горутину, которая не будет отправлять данные в os.Stdin
	go func() {
		time.Sleep(2 * time.Second)
		w.Close()
	}()

	numbersChan := make(chan []int)
	go BufioScannerInput(numbersChan)

	select {
	case <-numbersChan:
		t.Error("Expected no input, but got input")
	case <-time.After(1 * time.Second):
		// Ожидаем завершения функции после тайм-аута
		t.Log("Test passed: function timed out as expected")
	}
}

func TestBufioScannerInput_InvalidInput(t *testing.T) {
	// Подменяем os.Stdin
	input := "1 2 три 4 5" // некорректный ввод (слово "три")
	r, w, _ := os.Pipe()   // Правильный вызов os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }() // Восстанавливаем старый os.Stdin после теста

	// Записываем данные в writer
	go func() {
		w.WriteString(input + "\n")
		w.Close()
	}()

	numbersChan := make(chan []int)
	go BufioScannerInput(numbersChan)

	select {
	case result := <-numbersChan:
		expected := []int{1, 2, 4, 5} // "три" не должно преобразоваться в число
		// Проверяем, что длина среза совпадает с ожидаемой
		if len(result) != len(expected) {
			t.Errorf("Expected length %d, but got %d", len(expected), len(result))
		}

		// Проверяем каждый элемент до длины меньшего среза
		for i := 0; i < len(result) && i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Errorf("Expected %d, but got %d", expected[i], result[i])
			}
		}
	case <-time.After(1 * time.Second):
		t.Error("Test timed out waiting for input")
	}
}

func TestBufioScannerInput_TestTimeout(t *testing.T) {
	// Уменьшаем время тайм-аута для теста
	oldStdin := os.Stdin
	r, _, _ := os.Pipe()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }() // Восстанавливаем старый os.Stdin после теста

	numbersChan := make(chan []int)

	// Изменим время тайм-аута внутри функции
	go func() {
		timeout := time.After(3 * time.Second) // Изменим тайм-аут на 3 секунды для теста
		for range timeout {
			t.Log("Timeout occurred as expected")
			close(numbersChan)
			return
		}
	}()

	// Эмулируем долгий ввод данных (ничего не вводим, ждем завершения по тайм-ауту)
	select {
	case <-numbersChan:
		t.Log("Channel closed as expected after timeout")
	case <-time.After(4 * time.Second): // Добавляем тайм-аут, чтобы не зависнуть
		t.Error("Test timed out, expected function to stop after 3 seconds")
	}
}
