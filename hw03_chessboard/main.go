package main

import "fmt"

func main() {
	var size int
	fmt.Println("Введите размер доски:")

	fmt.Scanf("%d", &size)
	for stroka := 0; stroka < size; stroka++ {
		for kletka := 0; kletka < size; kletka++ {
			if (stroka+kletka)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
}
