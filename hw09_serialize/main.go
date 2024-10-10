// - Реализуйте структуру Book со следующими полями: ID, Title, Author, Year, Size, Rate (может быть дробным).
// - Реализуйте для нее интерфейсы Marshaller и Unmarshaller  из пакета json.
// - Составьте protobuf спецификацию для Book
// - Реализуйте для нее интерфейс Message из пакета proto.
// - Напишите фукции выполняющие сериализацию/десериализацию слайса объектов.
// - Напишите юнит тесты на реализованные функции;
package main

import "fmt"

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

type Marshaller interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
}

// func UnmarshallBook(book Marshaller) Unmarshaller {}

func main() {
	book := Book{
		ID:     1,
		Title:  "Go Programming Language",
		Author: "J. R. Tolkien",
		Year:   2009,
		Size:   420,
		Rate:   0.1,
	}
	fmt.Println(book)

}
