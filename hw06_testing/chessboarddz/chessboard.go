package chessboarddz

import (
	"errors"
)

func checksize(size int) error {
	if size <= 0 {
		return errors.New("размер должен быть больше нуля")
	}
	return nil
}

func paint(size int) string {
	if err := checksize(size); err != nil {
		return "Ошибка значения: " + err.Error()
	}

	var result string
	for stroka := 0; stroka < size; stroka++ {
		for kletka := 0; kletka < size; kletka++ {
			if (stroka+kletka)%2 == 0 {
				result += " "
			} else {
				result += "#"
			}
		}
		result += "\n"
	}
	return result
}
