package main

import "fmt"

func sort(pages *[]int) { // Сортируем массив
	for i := 0; i < len(*pages)-1; i++ {
		for j := i + 1; j < len(*pages); j++ {
			if (*pages)[i] > (*pages)[j] {
				(*pages)[i], (*pages)[j] = (*pages)[j], (*pages)[i]
			}
		}
	}
}

func binarSearch(needle int, pages []int) bool {
	low := 0               // Нулевое значение массива
	high := len(pages) - 1 // Максимальное значение массива
	for low <= high {
		mid := (low + high) / 2  // Среднее значение массива
		if pages[mid] < needle { // Если середина меньше искомого
			low = mid + 1
		} else {
			high = mid - 1
		}
		if low == len(pages) || pages[low] != needle {
			return false
		}
	}
	return true
}

func main() {
	pages := []int{12, 10, 3, 20}
	sort(&pages)
	fmt.Println(binarSearch(12, pages))
}
