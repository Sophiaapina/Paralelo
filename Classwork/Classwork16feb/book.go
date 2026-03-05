package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Println("Titles: ", b.Title, "\nAuthor: ", b.Author, "\nPages: ", b.Pages)
}
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

type TextBook struct {
	Book
	Editorial string
	Level     string
}

func (b *TextBook) PrintInfo() {
	fmt.Println("Title: ", b.Title, "\nAuthor :", b.Author, "\nPages :", b.Pages, "\nEditorial :", b.Editorial, "\nLevel :", b.Level)
}

func NewTextBook(title string, author string, pages int, editorial string,level string)*TextBook{
	return &TextBook{
		Book: book{title, author,pages},
	}
}
