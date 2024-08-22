package chessboarddz

import (
	"errors"
	"fmt"
)

var size int

func chessboard() int {
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
	return 0
}

func value() (int, error) {
	fmt.Println("Введите размер доски:")
	fmt.Scanf("%d", &size)
	// 	shape, ok := s.(Shape) // Приведение типа
	ok := size
	if ok <= 0 {
		return 0, errors.New("размер доски должен быть больше нуля")
	}
	return size, nil
}

func chessboarddz() {
	valueknow, _ := value()
	fmt.Println(valueknow)
	chess := chessboard()
	fmt.Println(chess)
}

//
