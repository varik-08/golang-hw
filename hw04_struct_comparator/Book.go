package main

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (book *Book) Id() int {
	return book.id
}

func (book *Book) SetId() int {
	return book.id
}

func (book *Book) Title() string {
	return book.title
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) Author() string {
	return book.author
}

func (book *Book) SetAuthor(author string) {
	book.author = author
}

func (book *Book) Year() int {
	return book.year
}

func (book *Book) SetYear(year int) {
	book.year = year
}

func (book *Book) Size() int {
	return book.size
}

func (book *Book) SetSize(size int) {
	book.size = size
}

func (book *Book) Rate() float32 {
	return book.rate
}

func (book *Book) SetRate(rate float32) {
	book.rate = rate
}
