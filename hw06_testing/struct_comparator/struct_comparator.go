package structcomparator

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
