package main

import (
	"errors"
	"fmt"
)

func checksize(size int) error {
	if size <= 0 {
		return errors.New("Размер должен быть больше нуля")
	}
	return nil
}

func paint(size int) {
	if err := checksize(size); err != nil {
		fmt.Println("Ошибка значения:", err)
		return
	}
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


