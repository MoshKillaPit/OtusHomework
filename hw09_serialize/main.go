package main

import (
	"encoding/json"
	"fmt"
	"log"

	protobook "github.com/MoshKillaPit/OtusHomework/hw09_serialize/proto"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b *Book) String() string {
	return fmt.Sprintf("Maket{ID: %v, Title: %v, Author: %v, Year: %v, Size: %v, Rate: %v}",
		b.ID, b.Title, b.Author, b.Year, b.Size, b.Rate)
}

func (b *Book) ProtoMessage() {}

func (b *Book) ProtoMarshal() ([]byte, error) {
	bookProto := &protobook.Book{
		ID:     int32(b.ID),
		Title:  b.Title,
		Author: b.Author,
		Year:   int32(b.Year),
		Size:   int32(b.Size),
		Rate:   b.Rate,
	}

	return proto.Marshal(bookProto)
}

func (b *Book) ProtoUnmarshal(data []byte) error {
	bookProto := &protobook.Book{}

	err := proto.Unmarshal(data, bookProto)
	if err != nil {
		return err
	}

	b.ID = int(bookProto.ID)
	b.Title = bookProto.Title
	b.Author = bookProto.Author
	b.Year = int(bookProto.Year)
	b.Size = int(bookProto.Size)
	b.Rate = bookProto.Rate

	return nil
}

func (b *Book) MarshalJSON() ([]byte, error) { // Делаем сериализацию
	type Alias Book               // Создаем псевдоним для избегания рекурсии
	return json.Marshal(&struct { // Оставляем структуру и представляем её в байтах
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error { // Делаем ансериализацию
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux) // Расшифровываем байты и передаём их в структуру
}

func main() {
	book := Book{ // Объявляем переменную и через неё передаём параметры в структуру книги
		ID:     1,
		Title:  "Book 1",
		Author: "Author 1",
		Year:   2001,
		Size:   2,
		Rate:   4.2,
	}
	jsonBytes, err := book.MarshalJSON() // Создаём новую переменую которая принимает слайс байтов из книги
	if err != nil {
		log.Fatalf("Ошибка сериализации JSON: %v", err)
	}
	fmt.Println(string(jsonBytes))

	protoBook := &protobook.Book{}
	err = json.Unmarshal(jsonBytes, protoBook)
	if err != nil {
		log.Fatalf("Ошибка десериализации JSON: %v", err)
	}
	fmt.Printf("Десериализованная книга: %+v\n", protoBook)

	protoBytes, err := book.ProtoMarshal()
	if err != nil {
		log.Fatalf("Ошибка сериализации PROTO: %v", err)
	}

	book.ProtoUnmarshal(protoBytes)

	fmt.Println(book)
}
