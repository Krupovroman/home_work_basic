package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(id int, title, author string, year, size int, rate float64) *Book {
	return &Book{id: id, title: title, author: author, year: year, size: size, rate: rate}
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) GetRate() float64 {
	return b.rate
}

// CompareMode представляет режим сравнения книг.
type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

type BookComparer struct {
	mode CompareMode
}

func NewBookComparer(mode CompareMode) *BookComparer {
	return &BookComparer{mode: mode}
}

func (bc *BookComparer) Compare(book1, book2 *Book) bool {
	switch bc.mode {
	case ByYear:
		return book1.year > book2.year
	case BySize:
		return book1.size > book2.size
	case ByRate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

func main() {
	book1 := NewBook(1, "Book 1", "Author 1", 1995, 400, 4.8)
	book2 := NewBook(2, "Book 2", "Author 2", 1990, 350, 4.9)

	comparerByYear := NewBookComparer(ByYear)
	fmt.Println("Book 2 is older than Book 1:", comparerByYear.Compare(book1, book2))

	comparerBySize := NewBookComparer(BySize)
	fmt.Println("Book 1 is bigger than Book 2:", comparerBySize.Compare(book1, book2))

	comparerByRate := NewBookComparer(ByRate)
	fmt.Println("Book 1 has higher rate than Book 2:", comparerByRate.Compare(book1, book2))
}
