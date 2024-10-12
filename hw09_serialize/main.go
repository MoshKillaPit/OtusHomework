package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

type Maket struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

type ProtoBook struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Size   int    `json:"size"`
	Rate   string `json:"rate"`
}

func (b *Maket) Reset() {
	b.ID = 0
	b.Title = ""
	b.Author = ""
	b.Year = 0
	b.Size = 0
	b.Rate = 0
}

func (b *Maket) String() string {
	return fmt.Sprintf("Maket{ID: %v, Title: %v, Author: %v, Year: %v, Size: %v, Rate: %v}", b.ID, b.Title, b.Author, b.Year, b.Size, b.Rate)
}

func (b *Maket) ProtoMessage() {

}
func (b *Maket) UnmarshalJSON(data []byte) error { // Делаем ансерилизацию для макета
	type Alias Maket
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}

func (b *Maket) ProtoMarshal() ([]byte, error) {
	return proto.Marshal(b)
}

func (b *Maket) ProtoUnmarshal(data []byte) error {
	return proto.Unmarshal(data, b)
}

func (b *Book) MarshalJSON() ([]byte, error) { // Делаем сериализацию
	type Alias Book               // Создаем псевдоним для избегания рекурсии
	return json.Marshal(&struct { // Оставляем структуру и представляем её в байтах
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}
func (b *Book) UnmarshalJSON(data []byte) error { //Делаем ансериализацию
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux) // Расшифровываем байты и передаём их в структуру
}

func (p *ProtoBook) ProtoUnmarshal(data []byte) error {
	return proto.Unmarshal(data, p)
}

func (p *ProtoBook) Reset() {
	p.ID = 0
	p.Title = ""
	p.Author = ""
	p.Year = 0
	p.Size = 0
}

func (p *ProtoBook) String() string {
	return fmt.Sprintf("Maket{ID: %v, Title: %v, Author: %v, Year: %v, Size: %v, Rate: %v}", p.ID, p.Title, p.Author, p.Year, p.Size, p.Rate)
}

func (p *ProtoBook) ProtoMessage() {

}

func main() {
	Books := Book{ // Объявляем переменную и через неё передаём параметры в структуру книги
		ID:     1,
		Title:  "Book 1",
		Author: "Author 1",
		Year:   2001,
		Size:   2,
		Rate:   4.2,
	}
	BookInfo, _ := Books.MarshalJSON() // Создаём новую переменую которая принимает слайс байтов из книги
	Maket := Maket{}                   // Объявляем новую переменную с привязкой с структуре макета
	Maket.UnmarshalJSON(BookInfo)      // Выполняем функцию преобразования байтов в структуру макета (Передача данных из книги (слайса байтов) в структуру макета)
	fmt.Println("Макет:", Maket)       // Посмотреть правильность данных переданных в макет
	// MAKETS DONE

	ProtoInfo, _ := Maket.ProtoMarshal()

	fmt.Println(ProtoInfo)

	ProtoBook := ProtoBook{}

	ProtoBook.ProtoUnmarshal(ProtoInfo)

	fmt.Println(ProtoBook)
}
