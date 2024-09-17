package main

import (
	"errors"
	"fmt"
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
	Height float64
	Width  float64
}
type triangle struct {
	Height float64
	Width  float64 // S = a * h / 2.
}

func (c circle) Calculation() float64 {
	return c.radius * pi
}

func (r rectangle) Calculation() float64 {
	return r.Height * r.Width
}

func (t triangle) Calculation() float64 {
	return t.Width / 2 * t.Height
}

func main() {
	Circle := circle{2.1}
	Rectangle := rectangle{Height: 12, Width: 0}
	Triangle := triangle{Height: 10, Width: 12}
	fmt.Println("Круг:", Circle.Calculation())
	fmt.Println("Прямоугольник:", Rectangle.Calculation())
	fmt.Println("Треугольник:", Triangle.Calculation())
	RandomCircle, err1 := calculateArea(Circle)
	if err1 != nil {
		fmt.Println("Ошибка", err1)
	} else {
		fmt.Println("Площадь круга:", RandomCircle)
	}
	RandomRectangle, err2 := calculateArea(Rectangle)
	if err2 != nil {
		fmt.Println("Ошибка", err2)
	} else {
		fmt.Println("Площадь прямоугольника:", RandomRectangle)
	}
	RandomTriangle, err3 := calculateArea(Triangle)
	if err3 != nil {
		fmt.Println("Ошибка", err3)
	} else {
		fmt.Println("Площадь круга:", RandomTriangle)
	}
}
