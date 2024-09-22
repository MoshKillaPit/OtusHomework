package shapes

import (
	"errors"
)

const pi float64 = 3.14

type Shape interface {
	Calculation() float64
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape) // Приведение типа
	if !ok {
		return 0, errors.New("переданный объект не реализует интерфейс Shape")
	}
	area := shape.Calculation()
	if area <= 0 {
		return 0, errors.New("площадь фигуры должна быть больше нуля")
	}
	return area, nil
}

type circle struct { // S = π × r в квадрате
	radius float64
}
type rectangle struct {
	height float64
	width  float64
}
type triangle struct {
	height float64
	width  float64 // S = a * h / 2.
}

func (c circle) Calculation() float64 {
	return c.radius * pi
}

func (r rectangle) Calculation() float64 {
	return r.height * r.width
}

func (t triangle) Calculation() float64 {
	return t.width / 2 * t.height
}
