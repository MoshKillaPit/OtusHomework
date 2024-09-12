package main

import "fmt"

type CompareMode int

const (
	CompareByYear CompareMode = iota
	CompareBySize
	CompareByRate
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (s *Book) SetID(id int) {
	s.id = id
}

func (s *Book) SetTitle(title string) {
	s.title = title
}

func (s *Book) SetAuthor(author string) {
	s.author = author
}

func (s *Book) SetYear(year int) {
	s.year = year
}

func (s *Book) SetSize(size int) {
	s.size = size
}

func (s *Book) SetRate(rate float64) {
	s.rate = rate
}

func (s *Book) GetID() int {
	return s.id
}

func (s *Book) GetTitle() string {
	return s.title
}

func (s *Book) GetAuthor() string {
	return s.author
}

func (s *Book) GetYear() int {
	return s.year
}

func (s *Book) GetSize() int {
	return s.size
}

func (s *Book) GetRate() float64 {
	return s.rate
}

type BookComparator struct {
	mode CompareMode
}

func (bc BookComparator) Compare(b1, b2 Book) bool {
	switch bc.mode {
	case CompareByYear:
		return b1.GetYear() == b2.GetYear()
	case CompareBySize:
		return b1.GetSize() == b2.GetSize()
	case CompareByRate:
		return b1.GetRate() == b2.GetRate()
	default:
		return false
	}
}

func NewComparator(mode CompareMode) BookComparator {
	return BookComparator{mode: mode}
}

func main() {
	Book1 := Book{
		id:     10,
		title:  "John Smith",
		author: "Dima",
		year:   2012,
		size:   20,
		rate:   5.4,
	}
	Book2 := Book{
		id:     20,
		title:  "Jane",
		author: "Rezanov",
		year:   2012,
		size:   10,
		rate:   2.4,
	}
	fmt.Println(Book1)

	Book1.SetID(1)
	fmt.Println(Book2.GetTitle())

	yearComparator := NewComparator(CompareByYear)
	fmt.Println(yearComparator.Compare(Book1, Book2))
	sizeComparator := NewComparator(CompareBySize)
	fmt.Println(sizeComparator.Compare(Book1, Book2))
	rateComparator := NewComparator(CompareByRate)
	fmt.Println(rateComparator.Compare(Book1, Book2))
}
